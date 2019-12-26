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
	Detallecargofamiliarsiradig       []Detallecargofamiliarsiradig       `json:"detallecargofamiliarsiradig"`
	Importegananciasotroempleosiradig []Importegananciasotroempleosiradig `json:"importegananciasotroempleosiradig"`
	Deducciondesgravacionsiradig      []Deducciondesgravacionsiradig      `json:"deducciondesgravacionsiradig"`
	Retencionpercepcionsiradig        []Retencionpercepcionsiradig        `json:"retencionpercepcionsiradig"`
	Beneficiosiradig                  []Beneficiosiradig                  `json:"beneficiosiradig"`
	Ajustesiradig                     []Ajustesiradig                     `json:"ajustesiradig"`
}
