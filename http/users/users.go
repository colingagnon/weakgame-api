//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package users

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	
	"net/http"
	
	"github.com/colingagnon/weakgame-api/lib"
	
	"strconv"

	"github.com/gorilla/mux"
	"github.com/colingagnon/weakgame-api/db/mysql/users"
	dbi "github.com/colingagnon/weakgame-api/dbi/users"
)

func GetLoginInfo(res http.ResponseWriter, req *http.Request) {
	var authId int;
	
	res, req, authId = lib.GetAuthUser(res, req)

	id := uint32(authId)
	
	if id > 0 {
		user, err := users.GetById(id)
		
		if err != nil {
			log.Println(err)
		}
		
		// Wipe hash from response
		user.Password = "";
		
		// Return user
		enc := json.NewEncoder(res)
		enc.Encode(user)
	} else {
		// TODO just throw 500 for now
		//d := map[string]string{"status":"fail", "message": "No session exists for this user"}
        //b := [1]map[string]string{d}
		//enc := json.NewEncoder(res)
		//enc.Encode(d)
		
		res.WriteHeader(500)
		fmt.Fprint(res, "No valid session found")
	}
}

func Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, OPTIONS")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Type, X-Authorization-Token")

	dec := json.NewDecoder(req.Body)
	row := &dbi.User{}
	dec.Decode(&row)
	
	row, err := users.GetByEmail(row.Email)
	if err != nil {
		log.Println(err)
	}
	
	// Check if rows is empty
	if row == nil {
		res.WriteHeader(500)
		fmt.Fprint(res, "Invalid credentials")
		return
	}
	
	
	// Remove password from response
	row.Password = ""
	
	// Generate token
	token := lib.GetAuthToken(int(row.Id));
	
	// Set header
	res.Header().Set("X-Authorization-Token", token)
	
	enc := json.NewEncoder(res)
	enc.Encode(row)
}

func Pre(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Authorization-Token")
}

func Post(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Authorization-Token")
	err := save(req)
	if err != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, err)
	}
}

// Returns the price of a revival
func GetRevive(res http.ResponseWriter, req *http.Request) {
	var authId int;
	
	res, req, authId = lib.GetAuthUser(res, req)

	id := uint32(authId);
	
	if id > 0 {
		user, _ := users.GetById(id)
		
		// Calculate fee
		d := map[string]float64{"fee":CalculateFee(float64(user.Experiencelevel))}
		
		enc := json.NewEncoder(res)
		enc.Encode(d)
	} else {
		res.WriteHeader(500)
		fmt.Fprint(res, "You must login to revive")
	    return;
	}
}

// Actually completes a rivival
func PostRevive(res http.ResponseWriter, req *http.Request) {
	var authId int;
	
	res, req, authId = lib.GetAuthUser(res, req)

	id := uint32(authId);
	
	if id > 0 {
		user, _ := users.GetById(id)
		
		if user.Hp != 0 {
			res.WriteHeader(500)
			fmt.Fprint(res, "You are not dead!")
		    return;
		}
		
		// Calculate fee
		fee := int32(CalculateFee(float64(user.Experiencelevel)));
		
		if user.Unicorns < fee {
			res.WriteHeader(500)
			fmt.Fprint(res, "You do not have enough unicorns to revive!")
		    return;
		}
		
		// Reduce number of unicorns
		user.Unicorns = user.Unicorns - fee
		
		user.Hp = CalculateHp(user.Experiencelevel)
		
		users.Save(user)
		
		// Hacky clear password hash from return
		user.Password = ""
		
		enc := json.NewEncoder(res)
		enc.Encode(user)
	} else {
		res.WriteHeader(500)
		fmt.Fprint(res, "You must login to revive")
	    return;
	}
}

func CalculateHp(level int32) int32 {
    return int32(2.00 + (float32(3.00 * level) * float32(1 + 0.06)))
}

func CalculateFee(level float64) float64 {
    return Round(1.00 + (float64(level) * 0.25));
}

func Round(f float64) float64 {
    return math.Floor(f + .5)
}

func List(res http.ResponseWriter, req *http.Request) {
	enc := json.NewEncoder(res)

	shouldFilter := false
	filterObject := &dbi.User{}
	
	/*
	req.ParseForm()
	if len(req.Form) > 0 {

		if _, ok := req.Form["id"]; ok {
			form_id := req.FormValue("id")
			shouldFilter = true
			i, _ := strconv.Atoi(form_id)
			parsedField := uint32(i)
			filterObject.Id = parsedField
		}

		if _, ok := req.Form["email"]; ok {
			form_email := req.FormValue("email")
			shouldFilter = true
			parsedField := form_email
			filterObject.Email = parsedField
		}

		if _, ok := req.Form["password"]; ok {
			form_password := req.FormValue("password")
			shouldFilter = true
			parsedField := form_password
			filterObject.Password = parsedField
		}

		if _, ok := req.Form["createdon"]; ok {
			form_createdon := req.FormValue("createdon")
			shouldFilter = true
			parsedField := form_createdon
			filterObject.Createdon = parsedField
		}

		if _, ok := req.Form["unicorns"]; ok {
			form_unicorns := req.FormValue("unicorns")
			shouldFilter = true
			i, _ := strconv.Atoi(form_unicorns)
			parsedField := int32(i)
			filterObject.Unicorns = parsedField
		}

		if _, ok := req.Form["hp"]; ok {
			form_hp := req.FormValue("hp")
			shouldFilter = true
			i, _ := strconv.Atoi(form_hp)
			parsedField := int32(i)
			filterObject.Hp = parsedField
		}

		if _, ok := req.Form["experiencelevel"]; ok {
			form_experiencelevel := req.FormValue("experiencelevel")
			shouldFilter = true
			i, _ := strconv.Atoi(form_experiencelevel)
			parsedField := int32(i)
			filterObject.Experiencelevel = parsedField
		}

		if _, ok := req.Form["experiencepoints"]; ok {
			form_experiencepoints := req.FormValue("experiencepoints")
			shouldFilter = true
			i, _ := strconv.Atoi(form_experiencepoints)
			parsedField := int32(i)
			filterObject.Experiencepoints = parsedField
		}

	}
	*/
	
	if shouldFilter {
		rows, err := users.Find(filterObject)
		if err != nil {
			log.Println(err)
		}
		enc.Encode(rows)
		return
	}

	rows, err := users.All()
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

	row, _ := users.GetById(id)
	enc.Encode(row)
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

	_, get_err := users.GetById(id)
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
	row := &dbi.User{}
	dec.Decode(&row)
	
	return users.Save(row)
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

	users.DeleteById(id)
}
