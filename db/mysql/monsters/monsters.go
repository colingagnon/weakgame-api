//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package monsters

import (
	"fmt"
	"math/rand"
	"time"
	"strings"

	dbi "github.com/cybermtl/weakgame-api/dbi/monsters"
	"is-a-dev.com/autoapi/lib"
	//"errors"
)

var DB lib.DB

//type MonsterCache struct{

//}

var cache = map[uint32]*dbi.Monster{}

func GetByRandom(levelLow int32, levelHigh int32) (*dbi.Monster, error) {
	row := &dbi.Monster{}
	
	//fmt.Println(levelLow);
	//fmt.Println(levelHigh);
	
	rows, err := DB.Query("SELECT id,url,monsterLevel,hp,accuracy,dodge,damageLow,damageHigh,experiencePoints FROM monsters WHERE monsterLevel BETWEEN ? AND ?",
		levelLow,
		levelHigh,
	)
	
	if err != nil {
		return nil, err
	}
	
	result := make([]*dbi.Monster, 0)
	for rows.Next() {
		r := &dbi.Monster{}
		rows.Scan(
			&r.Id,
			&r.Url,
			&r.Monsterlevel,
			&r.Hp,
			&r.Accuracy,
			&r.Dodge,
			&r.Damagelow,
			&r.Damagehigh,
			&r.Experiencepoints,
		)

		cache[r.Id] = r

		result = append(result, r)
	}
	
	if len(result) == 0 {
		return row, nil
	}
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	number := int(len(result) - 1)
	if number == 0{
		number = 1
	}
	
	randomMonsterIndex := r.Intn(number)
	
	return result[randomMonsterIndex], nil
}

func FindWithWhere(where string, params ...interface{}) ([]*dbi.Monster, error) {
	rows, err := DB.Query("SELECT id,url,monsterLevel,hp,accuracy,dodge,damageLow,damageHigh,experiencePoints FROM monsters "+where, params...)
	if err != nil {
		return nil, err
	}
	result := make([]*dbi.Monster, 0)
	for rows.Next() {
		r := &dbi.Monster{}
		rows.Scan(
			&r.Id,
			&r.Url,
			&r.Monsterlevel,
			&r.Hp,
			&r.Accuracy,
			&r.Dodge,
			&r.Damagelow,
			&r.Damagehigh,
			&r.Experiencepoints,
		)

		cache[r.Id] = r

		result = append(result, r)
	}
	return result, nil
}

func All() ([]*dbi.Monster, error) {
	return FindWithWhere("")
}

func GetById(id uint32) (*dbi.Monster, error) {

	row := &dbi.Monster{}
	err := DB.QueryRow("SELECT id,url,monsterLevel,hp,accuracy,dodge,damageLow,damageHigh,experiencePoints FROM monsters WHERE id = ?",
		id,
	).Scan(
		&row.Id,
		&row.Url,
		&row.Monsterlevel,
		&row.Hp,
		&row.Accuracy,
		&row.Dodge,
		&row.Damagelow,
		&row.Damagehigh,
		&row.Experiencepoints,
	)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func Find(monsters *dbi.Monster) ([]*dbi.Monster, error) {
	where := []string{}
	params := []interface{}{}

	if monsters.Id != 0 {
		where = append(where, "id = ?")
		params = append(params, monsters.Id)
	}

	if monsters.Url != "" {
		where = append(where, "url = ?")
		params = append(params, monsters.Url)
	}

	if monsters.Monsterlevel != 0 {
		where = append(where, "monsterLevel = ?")
		params = append(params, monsters.Monsterlevel)
	}

	if monsters.Hp != 0 {
		where = append(where, "hp = ?")
		params = append(params, monsters.Hp)
	}

	if monsters.Accuracy != 0 {
		where = append(where, "accuracy = ?")
		params = append(params, monsters.Accuracy)
	}

	if monsters.Dodge != 0 {
		where = append(where, "dodge = ?")
		params = append(params, monsters.Dodge)
	}

	if monsters.Damagelow != 0 {
		where = append(where, "damageLow = ?")
		params = append(params, monsters.Damagelow)
	}

	if monsters.Damagehigh != 0 {
		where = append(where, "damageHigh = ?")
		params = append(params, monsters.Damagehigh)
	}

	if monsters.Experiencepoints != 0 {
		where = append(where, "experiencePoints = ?")
		params = append(params, monsters.Experiencepoints)
	}

	var resultingwhere string
	if len(where) > 0 {
		resultingwhere = fmt.Sprintf("WHERE %s", strings.Join(where, " AND "))
	}
	return FindWithWhere(resultingwhere, params...)
}

func DeleteById(id uint32) error {
	//TODO: remove from cache.
	_, err := DB.Exec("DELETE FROM monsters WHERE id = ?",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func Save(row *dbi.Monster) error {
	if row.Monsterlevel == 0 {
		return lib.Error("Preconditions failed, Monsterlevel must be set.")
	}
	if row.Hp == 0 {
		return lib.Error("Preconditions failed, Hp must be set.")
	}
	if row.Accuracy == 0 {
		return lib.Error("Preconditions failed, Accuracy must be set.")
	}
	if row.Dodge == 0 {
		return lib.Error("Preconditions failed, Dodge must be set.")
	}
	if row.Damagelow == 0 {
		return lib.Error("Preconditions failed, Damagelow must be set.")
	}
	if row.Damagehigh == 0 {
		return lib.Error("Preconditions failed, Damagehigh must be set.")
	}
	if row.Experiencepoints == 0 {
		return lib.Error("Preconditions failed, Experiencepoints must be set.")
	}
	_, err := DB.Exec("INSERT monsters VALUES(?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE id = VALUES(id),url = VALUES(url),monsterlevel = VALUES(monsterlevel),hp = VALUES(hp),accuracy = VALUES(accuracy),dodge = VALUES(dodge),damagelow = VALUES(damagelow),damagehigh = VALUES(damagehigh),experiencepoints = VALUES(experiencepoints)",
		row.Id,
		row.Url,
		row.Monsterlevel,
		row.Hp,
		row.Accuracy,
		row.Dodge,
		row.Damagelow,
		row.Damagehigh,
		row.Experiencepoints,
	)
	if err != nil {
		return err
	}

	cache[row.Id] = row

	return nil
}
