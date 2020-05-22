package structLiquidacion

import (
	"time"

	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

//
type Liquidacion struct {
	structGormModel.GormModel
	Nombre                               string                    `json:"nombre"`
	Codigo                               string                    `json:"codigo"`
	Descripcion                          string                    `json:"descripcion"`
	Activo                               int                       `json:"activo"`
	Legajo                               *structLegajo.Legajo      `json:"legajo" gorm:"ForeignKey:Legajoid;association_foreignkey:ID;association_autoupdate:false;auto_preload:true;"`
	Legajoid                             *int                      `json:"legajoid" sql:"type:int REFERENCES Legajo(ID)" gorm:"not null;default:null"`
	Tipo                                 *Liquidaciontipo          `json:"tipo" gorm:"ForeignKey:Tipoid;association_foreignkey:ID;association_autoupdate:false;not null"` /*1-Mensual, 2-Primer Quincena, 3-Segunda Quincena, 4-Vacaciones, 5-SAC, 6-Liquidación Final*/
	Tipoid                               *int                      `json:"tipoid" sql:"type:int REFERENCES Liquidaciontipo(ID)" gorm:"not null"`
	Fecha                                time.Time                 `json:"fecha"`
	Fechaultimodepositoaportejubilatorio time.Time                 `json:"fechaultimodepositoaportejubilatorio"`
	Zonatrabajo                          string                    `json:"zonatrabajo"`
	Condicionpago                        *Liquidacioncondicionpago `json:"condicionpago" gorm:"ForeignKey:Condicionpagoid;association_foreignkey:ID;association_autoupdate:false;not null"` /*1-Cuenta Corriente, 2-Contado*/
	Condicionpagoid                      *int                      `json:"condicionpagoid" sql:"type:int REFERENCES Liquidacioncondicionpago(ID)" gorm:"not null"`
	Cuentabancoid                        *int                      `json:"cuentabancoid" gorm:"not null"`
	Cuentabanco                          *Banco                    `json:"cuentabanco"`
	Bancoaportejubilatorioid             *int                      `json:"bancoaportejubilatorioid" gorm:"not null;default:0"`
	Bancoaportejubilatorio               *Banco                    `json:"bancoaportejubilatorio"`
	Fechaperiododepositado               time.Time                 `json:"fechaperiododepositado"`
	Fechaperiodoliquidacion              time.Time                 `json:"fechaperiodoliquidacion"`
	Estacontabilizada                    bool                      `json:"estacontabilizada"`
	Asientomanualtransaccionid           int                       `json:"asientomanualtransaccionid"`
	Asientomanualnombre                  string                    `json:"asientomanualnombre"`
	Liquidacionitems                     []Liquidacionitem         `json:"liquidacionitems" gorm:"ForeignKey:Liquidacionid;association_foreignkey:ID"`
	Cantidaddiastrabajados               *int                      `json:"cantidaddiastrabajados" gorm:"not null;default:30"`
}
