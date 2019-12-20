package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Hijosiradig struct {
	structGormModel.GormModel
	Hijos             structLegajo.Hijo `json:"hijos" gorm:"ForeignKey:Hijoid;association_foreignkey:ID;association_autoupdate:false"`
	Hijoid            *int              `json:"hijoid" sql:"type:int REFERENCES Hijo(ID)"`
	Estaacargo        bool              `json:"estaacargo"`
	Residenteenelpais bool              `json:"residenteenelpais"`
	Obtuvoingresos    bool              `json:"obtuvoingresos"`
	Montoanual        *float64          `json:"montoanual"`
	Mesdesde          *time.Time        `json:"mesdesde"`
	Meshasta          *time.Time        `json:"meshasta"`
	Porcentaje        *float64          `json:"porcentaje"`
}
