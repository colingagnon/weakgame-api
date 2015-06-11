//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package products

type Producter interface {
	FindWithWhere(where string, params ...interface{}) ([]*Product, error)
	GetById(id uint32) (*Product, error)
	All() ([]*Product, error)
	Find(products *Product) ([]*Product, error)

	DeleteById(id uint32) error

	Save(row *Product) error
}

type Product struct {
	Id        uint32  `json:"id"`
	Name      string  `json:"name"`
	Tokens    int32   `json:"tokens"`
	Price     float32 `json:"price"`
	// TODO changed from byte[] array that was generated
	Available string  `json:"available"`
}

func New() *Product {
	return &Product{}
}
