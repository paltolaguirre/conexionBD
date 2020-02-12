package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Siradig struct {
	structGormModel.GormModel
	Legajo                            *structLegajo.Legajo                `json:"legajo" gorm:"ForeignKey:Legajoid;association_foreignkey:ID;association_autoupdate:false"`
	Legajoid                          *int                                `json:"legajoid" sql:"type:int REFERENCES Legajo(ID)" gorm:"not null"`
	Periodosiradig                    *time.Time                          `json:"periodosiradig"`
	Detallecargofamiliarsiradig       []Detallecargofamiliarsiradig       `json:"detallecargofamiliarsiradig" gorm:"ForeignKey:Siradigid;association_foreignkey:ID"`
	Importegananciasotroempleosiradig []Importegananciasotroempleosiradig `json:"importegananciasotroempleosiradig" gorm:"ForeignKey:Siradigid;association_foreignkey:ID"`
	Deducciondesgravacionsiradig      []Deducciondesgravacionsiradig      `json:"deducciondesgravacionsiradig" gorm:"ForeignKey:Siradigid;association_foreignkey:ID"`
	Retencionpercepcionsiradig        []Retencionpercepcionsiradig        `json:"retencionpercepcionsiradig" gorm:"ForeignKey:Siradigid;association_foreignkey:ID"`
	Beneficiosiradig                  []Beneficiosiradig                  `json:"beneficiosiradig" gorm:"ForeignKey:Siradigid;association_foreignkey:ID"`
	Ajustesiradig                     []Ajustesiradig                     `json:"ajustesiradig" gorm:"ForeignKey:Siradigid;association_foreignkey:ID"`
}
