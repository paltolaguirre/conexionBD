package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Retencionpercepcionsiradig struct {
	structGormModel.GormModel
	Siradigtipogrilla      Siradigtipogrilla    `json:"siradigtipogrilla" gorm:"ForeignKey:Siradigtipogrillaid;association_foreignkey:ID;association_autoupdate:false"`
	Siradigtipogrillaid    *int                 `json:"siradigtipogrillaid" sql:"type:int REFERENCES Siradigtipogrilla(ID)"`
	Siradigid              int                  `json:"siradigid"`
	Siradigotipoimpuesto   Siradigotipoimpuesto `json:"siradigtipoimpuesto" gorm:"ForeignKey:Siradigtipoimpuestoid;association_foreignkey:ID;association_autoupdate:false"`
	Siradigotipoimpuestoid *int                 `json:"siradigtipoimpuestoid" sql:"type:int REFERENCES Siradigtipoimpuesto(ID)"`
	Siradigtipooperacion   Siradigtipooperacion `json:"siradigtipooperacion" gorm:"ForeignKey:Siradigtipooperacionid;association_foreignkey:ID;association_autoupdate:false"`
	Siradigtipooperacionid *int                 `json:"siradigtipooperacionid" sql:"type:int REFERENCES Siradigtipooperacion(ID)"`
	Mes                    *time.Time           `json:"mes"`
	Importe                *float64             `json:"importe"  sql:"type:decimal(19,4);"`
	Descripcion            string               `json:"descripcion"`
}
