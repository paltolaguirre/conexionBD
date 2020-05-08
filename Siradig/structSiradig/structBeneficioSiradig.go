package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Beneficiosiradig struct {
	structGormModel.GormModel
	Siradigtipogrilla   Siradigtipogrilla `json:"siradigtipogrilla" gorm:"ForeignKey:Siradigtipogrillaid;association_foreignkey:ID;association_autoupdate:false"`
	Siradigtipogrillaid *int              `json:"siradigtipogrillaid" sql:"type:int REFERENCES Siradigtipogrilla(ID)"`
	Siradigid           int               `json:"siradigid"`
	Mesdesde            *time.Time        `json:"mesdesde"`
	Meshasta            *time.Time        `json:"meshasta"`
	Valor               bool              `json:"valor"`
}
