//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package fights

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"math/rand"
	"time"
	
	"github.com/colingagnon/weakgame-api/lib"
	
	"github.com/gorilla/mux"
	users "github.com/colingagnon/weakgame-api/db/mysql/users"
	monsters "github.com/colingagnon/weakgame-api/db/mysql/monsters"
	
	"github.com/colingagnon/weakgame-api/db/mysql/fights"
	dbi "github.com/colingagnon/weakgame-api/dbi/fights"	
)

// Create random number seed
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandom(res http.ResponseWriter, req *http.Request) {
	var authId int;
	
	res, req, authId = lib.GetAuthUser(res, req)

	id := uint32(authId);
	
	// Make sure we have a valid session
	if id > 0 {
	    user, _ := users.GetById(id)
	    
	    // Check if we already have a fight going
	    fightObject := &dbi.Fight{}
	    fightObject.Userid = id
	    
	    fightRows, err := fights.Find(fightObject)
		if err != nil {
			log.Println(err)
		}
		
		if len(fightRows) > 0 {
			enc := json.NewEncoder(res)
			enc.Encode(fightRows[0])
			return
		}
		
		// The level difference a user can encounter is a function of their level
		levelCur  := float32(user.Experiencelevel)
		levelDiff := float32(int32(levelCur * 0.20))
		
		levelLow := int32(levelCur - levelDiff) 
		levelHigh := int32(levelCur + levelDiff)
		
		// Now we can get our monster between levels
		monster, _ := monsters.GetByRandom(levelLow, levelHigh)
		
		if monster.Id == 0 {
			res.WriteHeader(500)
			fmt.Fprint(res, "Could not find a monster to fight")
		    return
		}
		
		// Finally we have our random monster, we need to create the fight now
		newFight := &dbi.Fight{}
		newFight.Userid = user.Id
		newFight.Monsterid = monster.Id
		newFight.Monstercurrenthp = monster.Hp
		newFight.Usercurrenthp = user.Hp
		
		errSave := fights.Save(newFight)
		if errSave != nil {
			fmt.Println(errSave);
		}
		
		// And now return the fight
		enc := json.NewEncoder(res)
		enc.Encode(newFight)
	} else {
		res.WriteHeader(500)
		fmt.Fprint(res, "You need to login before starting fights.")
	    return;
	}
}

func GetRound(res http.ResponseWriter, req *http.Request) {
	var authId int;
	
	res, req, authId = lib.GetAuthUser(res, req)

	id := uint32(authId);
	
	// Make sure we have a valid session
	if id > 0 {
	    user, _ := users.GetById(id)
	    
	    // Get current fight
	    fightObject := &dbi.Fight{}
	    fightObject.Userid = user.Id
	    
	    fightRows, err := fights.Find(fightObject)
		if err != nil {
			log.Println(err)
		}
		
		// Make sure we have a fight for this user and session
		if len(fightRows) == 0 {
			res.WriteHeader(500)
			fmt.Fprint(res, "You currently do not have a fight.")
	    	return;
		}
		
		fight := fightRows[0]
		
		// Get monster information
		// Now we can get our monster between levels
		monster, err := monsters.GetById(fight.Monsterid);
		if err != nil {
			log.Println(err)
			return
		}
		
		// User attacks
		var userBaseAccuracy, userBaseDodge float32 = 0.5, 0.05
		var acuracyLevelMod, dodgeLevelMod float32 = 0.02, 0.01
		var levelUpMod int32 = 50
		
		userAccuracy := userBaseAccuracy + (float32(user.Experiencelevel) * acuracyLevelMod)
		userDodge := userBaseDodge + (float32(user.Experiencelevel) * dodgeLevelMod)
		
		if getRandomHit(userAccuracy, monster.Dodge) {
			playerDamage := getPlayerDamage(user.Experiencelevel);
			fight.Monstercurrenthp -= playerDamage;
		}
		
		// Check if monster dead, award experience and delete fight
		if fight.Monstercurrenthp <= 0 {
			user.Experiencepoints += monster.Experiencepoints
			
			// Check for level up
			var level, totalExp int32 = 1, 0
			
			for {
				totalExp += levelUpMod * level * (1 + level)
				if totalExp > user.Experiencepoints {
					break;
				}
				level++
			}
			
			// Setup new level
			user.Experiencelevel = level;
			
			// Save user state
			users.Save(user)
			
			enc := json.NewEncoder(res)
			enc.Encode(fight)
			
			// Remove fight
			fights.DeleteById(uint32(fight.Id))
			if err != nil {
				fmt.Println("Failed to properly close fight");
			}
			
			return
		}
		
		// Monster attacks
		if getRandomHit(monster.Accuracy, userDodge) {
			monsterDamage := getMonsterDamage(monster.Damagelow, monster.Damagehigh);
			fight.Usercurrenthp -= monsterDamage;
		}
		
		// Check if monster dead, award experience and delete fight
		if fight.Usercurrenthp <= 0 {
			// TODO make better kill player
			user.Hp = 0
			
			// Save user state
			users.Save(user)
			
			enc := json.NewEncoder(res)
			enc.Encode(fight)
			
			// Remove fight
			err := fights.DeleteById(uint32(fight.Id))
			if err != nil {
				fmt.Println("Failed to properly close fight");
			}
			
			return
		}
		
		// Save fight state in case user goes away and logs in under different computer
		fights.Save(fight)
		
		// Always return fight status from this route
		enc := json.NewEncoder(res)
		enc.Encode(fight)
	} else {
		res.WriteHeader(500)
		fmt.Fprint(res, "You must login before continuing the fight.")
	    return;
	}
}

func getRandomHit(attackerAccuracy, defenderDodge float32) bool {
	isHit := false
	baseToHit := attackerAccuracy - defenderDodge
	
	// If base to hit is more than zero
	if baseToHit > 0 {
		result := r.Float32()
		
		if result < baseToHit {
			isHit = true;
		}
	}
	
	return isHit;
}

func getMonsterDamage(Damagelow, Damagehigh int32) int32 {
	return Damagelow + int32(r.Intn(int(Damagehigh - Damagelow + 1)));
}

func getPlayerDamage(level int32) int32 {
	// Calculation modifiers for player dmg
	lowMod, highMod, levelMod := 1.2, 2.0, 1.0
	
	// Base damage plus low and high modifiers, random
	userLevel := float64(level)
	baseDmg := int(levelMod * userLevel)
	lowDmg  := baseDmg + int(r.Intn(int(lowMod * userLevel)))
	highDmg := lowDmg + int(r.Intn(int(highMod * userLevel)))
	
	// Finally, randomly get damage from our range
	finalDmg := lowDmg + int(r.Intn((highDmg - lowDmg) + 1))
	
	return int32(finalDmg);
}	

func List(res http.ResponseWriter, req *http.Request) {
	enc := json.NewEncoder(res)

	shouldFilter := false
	filterObject := &dbi.Fight{}
	req.ParseForm()
	if len(req.Form) > 0 {

		if _, ok := req.Form["id"]; ok {
			form_id := req.FormValue("id")
			shouldFilter = true
			i, _ := strconv.Atoi(form_id)
			parsedField := uint32(i)
			filterObject.Id = parsedField
		}

		if _, ok := req.Form["monsterid"]; ok {
			form_monsterid := req.FormValue("monsterid")
			shouldFilter = true
			i, _ := strconv.Atoi(form_monsterid)
			parsedField := uint32(i)
			filterObject.Monsterid = parsedField
		}

		if _, ok := req.Form["userid"]; ok {
			form_userid := req.FormValue("userid")
			shouldFilter = true
			i, _ := strconv.Atoi(form_userid)
			parsedField := uint32(i)
			filterObject.Userid = parsedField
		}

		if _, ok := req.Form["usercurrenthp"]; ok {
			form_usercurrenthp := req.FormValue("usercurrenthp")
			shouldFilter = true
			i, _ := strconv.Atoi(form_usercurrenthp)
			parsedField := int32(i)
			filterObject.Usercurrenthp = parsedField
		}

		if _, ok := req.Form["monstercurrenthp"]; ok {
			form_monstercurrenthp := req.FormValue("monstercurrenthp")
			shouldFilter = true
			i, _ := strconv.Atoi(form_monstercurrenthp)
			parsedField := int32(i)
			filterObject.Monstercurrenthp = parsedField
		}

	}

	if shouldFilter {
		rows, err := fights.Find(filterObject)
		if err != nil {
			log.Println(err)
		}
		enc.Encode(rows)
		return
	}

	rows, err := fights.All()
	if err != nil {
		log.Println(err)
	}
	enc.Encode(rows)
}

func Get(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	enc := json.NewEncoder(res)

	param := vars["id"]

	i, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := uint32(i)

	row, _ := fights.GetById(id)
	enc.Encode(row)
}

func Post(res http.ResponseWriter, req *http.Request) {
	err := save(req)
	if err != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, err)
	}
}

func Put(res http.ResponseWriter, req *http.Request) {
	var err error
	vars := mux.Vars(req)

	param := vars["id"]

	i, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := uint32(i)

	_, get_err := fights.GetById(id)
	if get_err != nil {
		fmt.Println(get_err)
		fmt.Fprintln(res, get_err)
		return
	}
	err = save(req)
	if err != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, err)
	}
}

func save(req *http.Request) error {
	dec := json.NewDecoder(req.Body)
	row := &dbi.Fight{}
	dec.Decode(&row)
	return fights.Save(row)
}

func Delete(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	param := vars["id"]

	i, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := uint32(i)

	fights.DeleteById(id)
}
