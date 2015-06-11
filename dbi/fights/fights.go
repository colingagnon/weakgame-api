//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package fights

type Fighter interface {
	FindWithWhere(where string, params ...interface{}) ([]*Fight, error)
	GetById(id uint32) (*Fight, error)
	All() ([]*Fight, error)
	Find(fights *Fight) ([]*Fight, error)

	DeleteById(id uint32) error

	Save(row *Fight) error
}

type Fight struct {
	Id               uint32 `json:"id"`
	Monsterid        uint32 `json:"monsterId"`
	Userid           uint32 `json:"userId"`
	Usercurrenthp    int32  `json:"userCurrentHp"`
	Monstercurrenthp int32  `json:"monsterCurrentHp"`
}

func New() *Fight {
	return &Fight{}
}
