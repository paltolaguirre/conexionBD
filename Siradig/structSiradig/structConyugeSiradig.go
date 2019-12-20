package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Conyugesiradig struct {
	structGormModel.GormModel
	Conyuge           []structLegajo.Conyuge `json:"conyuge" gorm:"ForeignKey:Conyugeid;association_foreignkey:ID;association_autoupdate:false"`
	Conyugeid         *int                   `json:"conyugeid" sql:"type:int REFERENCES Conyuge(ID)"`
	Estaacargo        bool                   `json:"estaacargo"`
	Residenteenelpais bool                   `json:"residenteenelpais"`
	Obtuvoingresos    bool                   `json:"obtuvoingresos"`
	Montoanual        *float64               `json:"montoanual"`
	Mesdesde          *time.Time             `json:"mesdesde"`
	Meshasta          *time.Time             `json:"meshasta"`
}
