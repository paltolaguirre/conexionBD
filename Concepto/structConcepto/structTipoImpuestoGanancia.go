package structConcepto

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Tipoimpuestoganancia struct {
	structGormModel.GormModel
	Nombre                      string `json:"nombre"`
	Codigo                      string `json:"codigo"`
	Descripcion                 string `json:"descripcion"`
	Activo                      int    `json:"activo"`
	Aplicaimporteremunerativo   bool   `json:"aplicaimporteremunerativo"`
	Aplicaimportenoremunerativo bool   `json:"aplicaimportenoremunerativo"`
	Aplicadescuento             bool   `json:"aplicadescuento"`
	Aplicaretencion             bool   `json:"aplicaretencion"`
	Aplicaaportepatronal        bool   `json:"aplicaaportepatronal"`
}
