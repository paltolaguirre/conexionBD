package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Siradig struct {
	structGormModel.GormModel
	Legajo                              *structLegajo.Legajo                `json:"legajo" gorm:"ForeignKey:Legajoid;association_foreignkey:ID;association_autoupdate:false"`
	Legajoid                            *int                                `json:"legajoid" sql:"type:int REFERENCES Legajo(ID)" gorm:"not null"`
	Periodosiradig                      *time.Time                          `json:"periodosiradig"`
	Detallecargofamiliarsiradig         []Detallecargofamiliarsiradig       `json:"detallecargofamiliarsiradig"`
	Importegananciasotroempleosiradig   []Importegananciasotroempleosiradig `json:"importegananciasotroempleosiradig" gorm:"ForeignKey:Importegananciasotroempleosiradigid;association_foreignkey:ID;association_autoupdate:false"`
	Importegananciasotroempleosiradigid *int                                `json:"importegananciasotroempleosiradigid" sql:"type:int REFERENCES Importegananciasotroempleosiradig(ID)"`
	Deducciondesgravacionsiradig        *Deducciondesgravacionsiradig       `json:"deducciondesgravacionsiradig" gorm:"ForeignKey:deducciondesgravacionsiradigid;association_foreignkey:ID;association_autoupdate:false"`
	Deducciondesgravacionsiradigid      *int                                `json:"deducciondesgravacionsiradigid" sql:"type:int REFERENCES Deducciondesgravacionsiradig(ID)"`
	Retencionpercepcionsiradig          *Retencionpercepcionsiradig         `json:"retencionpercepcionsiradig" gorm:"ForeignKey:retencionpercepcionsiradigid;association_foreignkey:ID;association_autoupdate:false"`
	Retencionpercepcionsiradigid        *int                                `json:"retencionpercepcionsiradigid" sql:"type:int REFERENCES Retencionpercepcionsiradig(ID)"`
	Beneficiosiradig                    *Beneficiosiradig                   `json:"beneficiosiradig" gorm:"ForeignKey:beneficiosiradigid;association_foreignkey:ID;association_autoupdate:false"`
	Beneficiosiradigid                  *int                                `json:"beneficiosiradigid" sql:"type:int REFERENCES Beneficiosiradig(ID)"`
}

/*falta poner siradigid en cada struct y que todos sean []*/
