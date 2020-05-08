package structLiquidacion

import (
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Acumulador struct {
	structGormModel.GormModel
	Nombre            string   `json:"nombre"`
	Codigo            string   `json:"codigo"`
	Descripcion       string   `json:"descripcion"`
	Orden             int      `json:"orden"`
	Importe           *float64 `json:"importe" sql:"type:decimal(19,4);" gorm:"not null"`
	Tope              *float64 `json:"tope" sql:"type:decimal(19,4);"`
	Liquidacionitemid int      `json:"liquidacionitemid"`
	Esmostrable       bool     `json:"esmostrable"  gorm:"default:true"`
}
