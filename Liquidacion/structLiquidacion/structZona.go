package structLiquidacion

import (
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Zona struct {
	structGormModel.GormModel
	Nombre      string `json:"nombre"`
	Codigo      string `json:"codigo"`
	Descripcion string `json:"descripcion"`
	Reduccion   string `json:"reduccion"`
	Activo      int    `json:"activo"`
}
