package structTarea

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Tareaitem struct {
	structGormModel.GormModel
	Tareaid int    `json:"tareaid"`
	Json    string `json:"json" sql:"type:JSON;"`
	Estado  int    `json:"estado"`
	Detalle string `json:"detalle"`
}
