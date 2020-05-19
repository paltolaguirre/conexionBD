package structConcepto

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Conceptoafip struct {
	structGormModel.GormModel
	Nombre         string        `json:"nombre"`
	Codigo         string        `json:"codigo"`
	Descripcion    string        `json:"descripcion"`
	Activo         int           `json:"activo"`
	Tipoconcepto   *Tipoconcepto `json:"tipoconcepto" gorm:"ForeignKey:Tipoconceptoid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Tipoconceptoid *int          `json:"tipoconceptoid" gorm:"not null;default:0"`
}
