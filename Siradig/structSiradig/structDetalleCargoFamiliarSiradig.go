package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Detallecargofamiliarsiradig struct {
	structGormModel.GormModel
	Siradigtipogrilla   Siradigtipogrilla    `json:"siradigtipogrilla" gorm:"ForeignKey:Siradigtipogrillaid;association_foreignkey:ID;association_autoupdate:false"`
	Siradigtipogrillaid *int                 `json:"siradigtipogrillaid" sql:"type:int REFERENCES Siradigtipogrilla(ID)"`
	Siradigid           int                  `json:"siradigid"`
	Hijo                structLegajo.Hijo    `json:"hijos" gorm:"ForeignKey:Hijoid;association_foreignkey:ID;association_autoupdate:false"`
	Hijoid              *int                 `json:"hijoid" sql:"type:int REFERENCES Hijo(ID)"`
	Conyuge             structLegajo.Conyuge `json:"conyuge" gorm:"ForeignKey:Conyugeid;association_foreignkey:ID;association_autoupdate:false"`
	Conyugeid           *int                 `json:"conyugeid" sql:"type:int REFERENCES Conyuge(ID)"`
	Estaacargo          bool                 `json:"estaacargo"`
	Residenteenelpais   bool                 `json:"residenteenelpais"`
	Obtuvoingresos      bool                 `json:"obtuvoingresos"`
	Montoanual          *float64             `json:"montoanual"`
	Mesdesde            *time.Time           `json:"mesdesde"`
	Meshasta            *time.Time           `json:"meshasta"`
	Porcentaje          *float64             `json:"porcentaje"`
}
