package structConcepto

import (
	"github.com/xubiosueldos/conexionBD/Function/structFunction"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

//NO ESTOY USANDO EL MODEL DE GORM PORQUE SOLO USA UNSIGNED INTS ENTONCES NO PUEDO USAR ID NEGATIVOS
type Concepto struct {
	structGormModel.GormModel
	Nombre                  *string                  `json:"nombre" gorm:"not null"`
	Codigo                  string                   `json:"codigo"`
	Descripcion             string                   `json:"descripcion"`
	Activo                  int                      `json:"activo"`
	Tipo                    string                   `json:"tipo"`
	CuentaContable          *int                     `json:"cuentacontableid" gorm:"not null"`
	Cuentacontable          *Cuenta                  `json:"cuentacontable"`
	Esimprimible            bool                     `json:"esimprimible"`
	Tipoconcepto            *Tipoconcepto            `json:"tipoconcepto" gorm:"ForeignKey:Tipoconceptoid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Tipoconceptoid          *int                     `json:"tipoconceptoid" gorm:"not null;default:0"`
	Esnovedad               bool                     `json:"esnovedad" gorm:"default:false"`
	Porcentaje              *float64                 `json:"porcentaje" sql:"type:decimal(19,4);"`
	Tipodecalculo           *Tipodecalculo           `json:"tipodecalculo" gorm:"ForeignKey:Tipodecalculoid;association_foreignkey:ID;association_autoupdate:false"`
	Tipodecalculoid         *int                     `json:"tipodecalculoid"`
	Prorrateo               bool                     `json:"prorrateo"`
	Basesac                 bool                     `json:"basesac"`
	Tipoimpuestoganancias   *Tipoimpuestoganancias   `json:"tipoimpuestoganancias" gorm:"ForeignKey:Tipoimpuestogananciasid;association_foreignkey:ID;association_autoupdate:false"`
	Tipoimpuestogananciasid *int                     `json:"tipoimpuestogananciasid"`
	Eseditable              bool                     `json:"eseditable"`
	Tipocalculoautomatico   *Tipocalculoautomatico   `json:"tipocalculoautomatico" gorm:"ForeignKey:Tipocalculoautomaticoid;association_foreignkey:ID;association_autoupdate:false"`
	Tipocalculoautomaticoid *int                     `json:"tipocalculoautomaticoid"`
	Formula                 *structFunction.Function `json:"formula" gorm:"ForeignKey:formulanombre;association_foreignkey:Name;association_autoupdate:false"`
	Formulanombre           *string                  `json:"formulanombre"`
	Esremvariable           bool                     `json:"esremvariable"`
	Conceptoafip            *Conceptoafip            `json:"conceptoafip" gorm:"ForeignKey:conceptoafipid;association_foreignkey:ID;association_autoupdate:false"`
	Conceptoafipid          *int                     `json:"conceptoafipid"`
}
