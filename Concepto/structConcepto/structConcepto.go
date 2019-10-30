package structConcepto

import "github.com/xubiosueldos/conexionBD/structGormModel"

//NO ESTOY USANDO EL MODEL DE GORM PORQUE SOLO USA UNSIGNED INTS ENTONCES NO PUEDO USAR ID NEGATIVOS
type Concepto struct {
	structGormModel.GormModel
	Nombre         *string       `json:"nombre" gorm:"not null"`
	Codigo         string        `json:"codigo"`
	Descripcion    string        `json:"descripcion"`
	Activo         int           `json:"activo"`
	Tipo           string        `json:"tipo"`
	CuentaContable *int          `json:"cuentacontableid" gorm:"not null"`
	Cuentacontable *Cuenta       `json:"cuentacontable"`
	Esimprimible   bool          `json:"esimprimible"`
	Tipoconcepto   *Tipoconcepto `json:"tipoconcepto" gorm:"ForeignKey:Tipoconceptoid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Tipoconceptoid *int          `json:"tipoconceptoid" gorm:"not null;default:0"`
	Esnovedad      bool          `json:"esnovedad" gorm:"default:false"`
}
