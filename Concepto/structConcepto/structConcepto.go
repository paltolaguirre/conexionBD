package structConcepto

import (
	"github.com/xubiosueldos/conexionBD/Function/structFunction"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

//NO ESTOY USANDO EL MODEL DE GORM PORQUE SOLO USA UNSIGNED INTS ENTONCES NO PUEDO USAR ID NEGATIVOS
type Concepto struct {
	structGormModel.GormModel
	Nombre                        *string                  `json:"nombre" gorm:"not null"`
	Codigo                        string                   `json:"codigo"`
	Descripcion                   string                   `json:"descripcion"`
	Activo                        int                      `json:"activo"`
	Tipo                          string                   `json:"tipo"`
	CuentaContable                *int                     `json:"cuentacontableid" gorm:"not null"`
	Cuentacontable                *Cuenta                  `json:"cuentacontable"`
	Esimprimible                  bool                     `json:"esimprimible"`
	Tipoconcepto                  *Tipoconcepto            `json:"tipoconcepto" gorm:"ForeignKey:Tipoconceptoid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Tipoconceptoid                *int                     `json:"tipoconceptoid" gorm:"not null;default:0"`
	Esnovedad                     bool                     `json:"esnovedad" gorm:"default:false"`
	Porcentaje                    *float64                 `json:"porcentaje" sql:"type:decimal(19,4);"`
	Tipodecalculo                 *Tipodecalculo           `json:"tipodecalculo" gorm:"ForeignKey:Tipodecalculoid;association_foreignkey:ID;association_autoupdate:false"`
	Tipodecalculoid               *int                     `json:"tipodecalculoid"`
	Prorrateo                     bool                     `json:"prorrateo"`
	Basesac                       bool                     `json:"basesac"`
	Tipoimpuestoganancias         *Tipoimpuestoganancias   `json:"tipoimpuestoganancias" gorm:"ForeignKey:Tipoimpuestogananciasid;association_foreignkey:ID;association_autoupdate:false"`
	Tipoimpuestogananciasid       *int                     `json:"tipoimpuestogananciasid"`
	Eseditable                    bool                     `json:"eseditable"`
	Tipocalculoautomatico         *Tipocalculoautomatico   `json:"tipocalculoautomatico" gorm:"ForeignKey:Tipocalculoautomaticoid;association_foreignkey:ID;association_autoupdate:false"`
	Tipocalculoautomaticoid       *int                     `json:"tipocalculoautomaticoid"`
	Formula                       *structFunction.Function `json:"formula" gorm:"ForeignKey:formulanombre;association_foreignkey:Name;association_autoupdate:false"`
	Formulanombre                 *string                  `json:"formulanombre"`
	Esremvariable                 bool                     `json:"esremvariable"`
	Conceptoafip                  *Conceptoafip            `json:"conceptoafip" gorm:"ForeignKey:conceptoafipid;association_foreignkey:ID;association_autoupdate:false"`
	Conceptoafipid                *int                     `json:"conceptoafipid"`
	Codigointerno                 *int                     `json:"codigointerno" gorm:"auto_increment;size:9"`
	Marcarepeticion               bool                     `json:"marcarepeticion" gorm:"default:false"`
	Aportesipa                    bool                     `json:"aportesipa" gorm:"default:false"`
	Contribucionsipa              bool                     `json:"contribucionsipa" gorm:"default:false"`
	Aportesinssjyp                bool                     `json:"aportesinssjyp" gorm:"default:false"`
	Contribucionesinssjyp         bool                     `json:"contribucionesinssjyp" gorm:"default:false"`
	Aportesobrasocial             bool                     `json:"aportesobrasocial" gorm:"default:false"`
	Contribucionesobrasocial      bool                     `json:"contribucionesobrasocial" gorm:"default:false"`
	Aportesfondosolidario         bool                     `json:"aportesfondosolidario" gorm:"default:false"`
	Contribucionesfondosolidario  bool                     `json:"contribucionesfondosolidario" gorm:"default:false"`
	Aportesrenatea                bool                     `json:"aportesrenatea" gorm:"default:false"`
	Contribucionesrenatea         bool                     `json:"contribucionesrenatea" gorm:"default:false"`
	Asignacionesfamiliares        bool                     `json:"asignacionesfamiliares" gorm:"default:false"`
	Contribucionesfondonacional   bool                     `json:"contribucionesfondonacional" gorm:"default:false"`
	Contribucionesleyriesgo       bool                     `json:"contribucionesleyriesgo" gorm:"default:false"`
	Aportesregimenesdiferenciales bool                     `json:"aportesregimenesdiferenciales" gorm:"default:false"`
	Aportesregimenesespeciales    bool                     `json:"aportesregimenesespeciales" gorm:"default:false"`
	Cuentacontablepasivoid        *int                     `json:"cuentacontablepasivoid"`
	Cuentacontablepasivo          *Cuenta                  `json:"cuentacontablepasivo"`
	Esganancias                   bool                     `json:"esganancias" gorm:"default:false"`
}
