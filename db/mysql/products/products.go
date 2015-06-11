//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package products

import (
	"fmt"
	"strings"

	dbi "github.com/colingagnon/weakgame-api/dbi/products"
	"is-a-dev.com/autoapi/lib"
	//"errors"
)

var DB lib.DB

//type ProductCache struct{

//}

var cache = map[uint32]*dbi.Product{}

func FindWithWhere(where string, params ...interface{}) ([]*dbi.Product, error) {
	rows, err := DB.Query("SELECT id,name,tokens,price,available FROM products "+where, params...)
	if err != nil {
		return nil, err
	}
	result := make([]*dbi.Product, 0)
	for rows.Next() {
		r := &dbi.Product{}
		rows.Scan(
			&r.Id,
			&r.Name,
			&r.Tokens,
			&r.Price,
			&r.Available,
		)

		cache[r.Id] = r

		result = append(result, r)
	}
	return result, nil
}

func All() ([]*dbi.Product, error) {
	return FindWithWhere("")
}

func GetById(id uint32) (*dbi.Product, error) {

	row := &dbi.Product{}
	err := DB.QueryRow("SELECT id,name,tokens,price,available FROM products WHERE id = ?",
		id,
	).Scan(
		&row.Id,
		&row.Name,
		&row.Tokens,
		&row.Price,
		&row.Available,
	)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func Find(products *dbi.Product) ([]*dbi.Product, error) {
	where := []string{}
	params := []interface{}{}

	if products.Id != 0 {
		where = append(where, "id = ?")
		params = append(params, products.Id)
	}

	if products.Name != "" {
		where = append(where, "name = ?")
		params = append(params, products.Name)
	}

	if products.Tokens != 0 {
		where = append(where, "tokens = ?")
		params = append(params, products.Tokens)
	}

	if products.Price != 0 {
		where = append(where, "price = ?")
		params = append(params, products.Price)
	}

	// TODO changed from byte[] array that was generated
	if products.Available != "" {
		where = append(where, "available = ?")
		params = append(params, products.Available)
	}

	var resultingwhere string
	if len(where) > 0 {
		resultingwhere = fmt.Sprintf("WHERE %s", strings.Join(where, " AND "))
	}
	return FindWithWhere(resultingwhere, params...)
}

func DeleteById(id uint32) error {
	//TODO: remove from cache.
	_, err := DB.Exec("DELETE FROM products WHERE id = ?",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func Save(row *dbi.Product) error {
	if row.Tokens == 0 {
		return lib.Error("Preconditions failed, Tokens must be set.")
	}
	if row.Price == 0 {
		return lib.Error("Preconditions failed, Price must be set.")
	}
	fmt.Println()
	_, err := DB.Exec("INSERT products VALUES(?,?,?,?,?) ON DUPLICATE KEY UPDATE id = VALUES(id),name = VALUES(name),tokens = VALUES(tokens),price = VALUES(price),available = VALUES(available)",
		row.Id,
		row.Name,
		row.Tokens,
		row.Price,
		row.Available,
	)
	if err != nil {
		return err
	}

	cache[row.Id] = row

	return nil
}
