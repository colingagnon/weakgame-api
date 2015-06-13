//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package users

import (
	"fmt"
	"strings"
	"time"
	"crypto/sha256"
	"io"
	
	dbi "github.com/cybermtl/weakgame-api/dbi/users"
	"is-a-dev.com/autoapi/lib"
	//"errors"
)

var DB lib.DB

//type UserCache struct{

//}

var cache = map[uint32]*dbi.User{}

func FindWithWhere(where string, params ...interface{}) ([]*dbi.User, error) {
	//fmt.Println(params)
	//fmt.Println("SELECT id,email,password,createdOn,unicorns,hp,experienceLevel,experiencePoints FROM users "+where)
	rows, err := DB.Query("SELECT id,email,password,createdOn,unicorns,hp,experienceLevel,experiencePoints FROM users "+where, params...)
	if err != nil {
		return nil, err
	}
	result := make([]*dbi.User, 0)
	for rows.Next() {
		r := &dbi.User{}
		rows.Scan(
			&r.Id,
			&r.Email,
			&r.Password,
			&r.Createdon,
			&r.Unicorns,
			&r.Hp,
			&r.Experiencelevel,
			&r.Experiencepoints,
		)

		cache[r.Id] = r

		result = append(result, r)
	}
	return result, nil
}

func All() ([]*dbi.User, error) {
	return FindWithWhere("")
}

func GetById(id uint32) (*dbi.User, error) {

	row := &dbi.User{}
	err := DB.QueryRow("SELECT id,email,password,createdOn,unicorns,hp,experienceLevel,experiencePoints FROM users WHERE id = ?",
		id,
	).Scan(
		&row.Id,
		&row.Email,
		&row.Password,
		&row.Createdon,
		&row.Unicorns,
		&row.Hp,
		&row.Experiencelevel,
		&row.Experiencepoints,
	)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func GetByEmail(email string) (*dbi.User, error) {
	
	row := &dbi.User{}
	err := DB.QueryRow("SELECT id,email,password,createdOn,unicorns,hp,experienceLevel,experiencePoints FROM users WHERE email = ?",
		email,
	).Scan(
		&row.Id,
		&row.Email,
		&row.Password,
		&row.Createdon,
		&row.Unicorns,
		&row.Hp,
		&row.Experiencelevel,
		&row.Experiencepoints,
	)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func Find(users *dbi.User) ([]*dbi.User, error) {
	where := []string{}
	params := []interface{}{}

	if users.Id != 0 {
		where = append(where, "id = ?")
		params = append(params, users.Id)
	}

	if users.Email != "" {
		where = append(where, "email = ?")
		params = append(params, users.Email)
	}

	if users.Password != "" {
		where = append(where, "password = ?")
		params = append(params, hashPassword(users.Password))
	}

	if !users.Createdon.IsZero() {
		where = append(where, "createdOn = ?")
		params = append(params, users.Createdon)
	}

	if users.Unicorns != 0 {
		where = append(where, "unicorns = ?")
		params = append(params, users.Unicorns)
	}

	if users.Hp != 0 {
		where = append(where, "hp = ?")
		params = append(params, users.Hp)
	}

	if users.Experiencelevel != 0 {
		where = append(where, "experienceLevel = ?")
		params = append(params, users.Experiencelevel)
	}

	if users.Experiencepoints != 0 {
		where = append(where, "experiencePoints = ?")
		params = append(params, users.Experiencepoints)
	}

	var resultingwhere string
	if len(where) > 0 {
		resultingwhere = fmt.Sprintf("WHERE %s", strings.Join(where, " AND "))
	}
	
	return FindWithWhere(resultingwhere, params...)
}

func DeleteById(id uint32) error {
	//TODO: remove from cache.
	_, err := DB.Exec("DELETE FROM users WHERE id = ?",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func hashPassword (password string) string {
	h := sha256.New()
	io.WriteString(h, password)
	password = fmt.Sprintf("%x", h.Sum(nil))
	
	return password;
}

func Save(row *dbi.User) error {
	if row.Email == "" {
		return lib.Error("Preconditions failed, Email must be set.")
	}
	if row.Password == "" {
		return lib.Error("Preconditions failed, Password must be set.")
	}
	
	existing, _ := GetByEmail(row.Email)
	if existing != nil && row.Id == 0 {
		return lib.Error("This user already exists")
	}

	if row.Id == 0 {
		row.Password = hashPassword(row.Password)
		row.Experiencelevel = 1
		row.Createdon = time.Now()
		row.Hp = 5
		row.Unicorns = 5
	}
	
	// TODO weird bug if you reduce the number of parameters, it doesn't seem to work
	_, err := DB.Exec("INSERT users VALUES(?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE id = VALUES(id),email = VALUES(email),password = VALUES(password),createdon = VALUES(createdon),unicorns = VALUES(unicorns),hp = VALUES(hp),experiencelevel = VALUES(experiencelevel),experiencepoints = VALUES(experiencepoints)",
		row.Id,
		row.Email,
		row.Password,
		row.Createdon,
		row.Unicorns,
		row.Hp,
		row.Experiencelevel,
		row.Experiencepoints,
	)
	if err != nil {
		return err
	}

	cache[row.Id] = row

	return nil
}
