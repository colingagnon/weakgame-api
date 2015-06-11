//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package monsters

type Monsterer interface {
	FindWithWhere(where string, params ...interface{}) ([]*Monster, error)
	GetById(id uint32) (*Monster, error)
	All() ([]*Monster, error)
	Find(monsters *Monster) ([]*Monster, error)

	DeleteById(id uint32) error

	Save(row *Monster) error
}

type Monster struct {
	Id               uint32  `json:"id"`
	Url              string  `json:"url"`
	Monsterlevel     int32   `json:"monsterLevel"`
	Hp               int32   `json:"hp"`
	Accuracy         float32 `json:"accuracy"`
	Dodge            float32 `json:"dodge"`
	Damagelow        int32   `json:"damageLow"`
	Damagehigh       int32   `json:"damageHigh"`
	Experiencepoints int32   `json:"experiencePoints"`
}

func New() *Monster {
	return &Monster{}
}
