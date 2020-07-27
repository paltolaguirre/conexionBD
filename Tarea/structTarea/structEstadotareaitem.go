package structTarea

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Estadotareaitem struct {
	structGormModel.GormModel
	Nombre      string `json:"nombre"`
	Codigo      string `json:"codigo"`
	Descripcion string `json:"descripcion"`
	Activo      int    `json:"activo"`
}
