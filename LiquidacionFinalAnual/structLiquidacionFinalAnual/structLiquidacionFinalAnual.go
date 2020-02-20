package structLiquidacionFinalAnual

import (
	"time"

	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Liquidacionfinalanual struct {
	structGormModel.GormModel
	Tipopresentacion      *Tipopresentacionliquidacionfinalanual `json:"tipopresentacion" gorm:"ForeignKey:Tipopresentacionid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Tipopresentacionid    *int                                   `json:"tipopresentacionid" sql:"type:int REFERENCES Tipopresentacionliquidacionfinalanual(ID)"`
	Periodofiscal         *time.Time                             `json:"periodofiscal"`
	Legajo                *structLegajo.Legajo                   `json:"legajo" gorm:"ForeignKey:Legajoid;association_foreignkey:ID;association_autoupdate:false;auto_preload:true;"`
	Legajoid              *int                                   `json:"legajoid" sql:"type:int REFERENCES Legajo(ID)"`
	Remuneraciones        *float64                               `json:"remuneraciones"`
	Deduccionesgenerales  *float64                               `json:"deduccionesgenerales"`
	Deduccionespersonales *float64                               `json:"deduccionespersonales"`
	Impuestodeterminado   *float64                               `json:"impuestodeterminado"`
}
