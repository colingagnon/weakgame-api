//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package fights

import (
	"fmt"
	"strings"
	
	dbi "github.com/colingagnon/weakgame-api/dbi/fights"
	"is-a-dev.com/autoapi/lib"
	//"errors"
)

var DB lib.DB

var cache = map[uint32]*dbi.Fight{}

func FindWithWhere(where string, params ...interface{}) ([]*dbi.Fight, error) {
	rows, err := DB.Query("SELECT id,monsterId,userId,userCurrentHp,monsterCurrentHp FROM fights "+where, params...)
	if err != nil {
		return nil, err
	}
	result := make([]*dbi.Fight, 0)
	for rows.Next() {
		r := &dbi.Fight{}
		rows.Scan(
			&r.Id,
			&r.Monsterid,
			&r.Userid,
			&r.Usercurrenthp,
			&r.Monstercurrenthp,
		)

		cache[r.Id] = r

		result = append(result, r)
	}
	
	return result, nil
}

func All() ([]*dbi.Fight, error) {
	return FindWithWhere("")
}

func GetById(id uint32) (*dbi.Fight, error) {

	row := &dbi.Fight{}
	err := DB.QueryRow("SELECT id,monsterId,userId,userCurrentHp,monsterCurrentHp FROM fights WHERE id = ?",
		id,
	).Scan(
		&row.Id,
		&row.Monsterid,
		&row.Userid,
		&row.Usercurrenthp,
		&row.Monstercurrenthp,
	)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func Find(fights *dbi.Fight) ([]*dbi.Fight, error) {
	where := []string{}
	params := []interface{}{}

	if fights.Id != 0 {
		where = append(where, "id = ?")
		params = append(params, fights.Id)
	}

	if fights.Monsterid != 0 {
		where = append(where, "monsterId = ?")
		params = append(params, fights.Monsterid)
	}

	if fights.Userid != 0 {
		where = append(where, "userId = ?")
		params = append(params, fights.Userid)
	}

	if fights.Usercurrenthp != 0 {
		where = append(where, "userCurrentHp = ?")
		params = append(params, fights.Usercurrenthp)
	}

	if fights.Monstercurrenthp != 0 {
		where = append(where, "monsterCurrentHp = ?")
		params = append(params, fights.Monstercurrenthp)
	}

	var resultingwhere string
	if len(where) > 0 {
		resultingwhere = fmt.Sprintf("WHERE %s", strings.Join(where, " AND "))
	}
	return FindWithWhere(resultingwhere, params...)
}

func DeleteById(id uint32) error {	
	//TODO: remove from cache
	_, err := DB.Query("DELETE FROM fights WHERE id = ?",
		id,
	)
	
	if err != nil {
		return err
	}
	
	return nil
}

func Save(row *dbi.Fight) error {
	if row.Monsterid == 0 {
		return lib.Error("Preconditions failed, Monsterid must be set.")
	}
	if row.Userid == 0 {
		return lib.Error("Preconditions failed, Userid must be set.")
	}
	/*
	if row.Usercurrenthp == 0 {
		return lib.Error("Preconditions failed, Usercurrenthp must be set.")
	}
	if row.Monstercurrenthp == 0 {
		return lib.Error("Preconditions failed, Monstercurrenthp must be set.")
	}
	*/
	_, err := DB.Exec("INSERT fights VALUES(?,?,?,?,?) ON DUPLICATE KEY UPDATE id = VALUES(id),monsterid = VALUES(monsterid),userid = VALUES(userid),usercurrenthp = VALUES(usercurrenthp),monstercurrenthp = VALUES(monstercurrenthp)",
		row.Id,
		row.Monsterid,
		row.Userid,
		row.Usercurrenthp,
		row.Monstercurrenthp,
	)
	if err != nil {
		return err
	}

	cache[row.Id] = row

	return nil
}
