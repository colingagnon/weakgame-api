//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package db

import (
	"database/sql"
	"errors"
	"fmt"

	"is-a-dev.com/autoapi/lib"
)

//Checksum is an autoapi-generated checksum of the state of the database, at time of generation.
func Checksum() string {
	return "da6cc18b40df664211e6c2c5bbb59fc0"
}

//ValidateChecksum compares the checksum generated by Autoapi to the current state of the db,
//returning an error if they don't match.
func ValidateChecksum(db *sql.DB, dbName string) error {
	b, err := lib.DatabaseChecksum(db, dbName)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%x", b) != Checksum() {
		fmt.Println(fmt.Sprintf("%x", b))
		fmt.Println(Checksum())
		return ErrBadDatabaseChecksum
	}
	return nil
}

//MustValidateChecksum compares the checksum against the database, and panics if they don't match.
//Useful when you absolutely don't want to run the software against a non-matching version of the db.
func MustValidateChecksum(db *sql.DB, dbName string) {
	if err := ValidateChecksum(db, dbName); err != nil {
		panic(err)
	}
}

var ErrBadDatabaseChecksum = errors.New("The code doesn't match the database's structure.")
