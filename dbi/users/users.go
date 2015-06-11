//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package users

import "time"

type Userer interface {
	FindWithWhere(where string, params ...interface{}) ([]*User, error)
	GetById(id uint32) (*User, error)
	All() ([]*User, error)
	Find(users *User) ([]*User, error)

	DeleteById(id uint32) error

	Save(row *User) error
}

type User struct {
	Id               uint32    `json:"id"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	Createdon        time.Time `json:"createdOn"`
	Unicorns         int32     `json:"unicorns"`
	Hp               int32     `json:"hp"`
	Experiencelevel  int32     `json:"experienceLevel"`
	Experiencepoints int32     `json:"experiencePoints"`
}

func New() *User {
	return &User{}
}
