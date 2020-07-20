package structTarea

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Tarea struct {
	structGormModel.GormModel
	Handler    string      `json:"handler"`
	TareaItems []Tareaitem `json:"tareaitems" gorm:"ForeignKey:Tareaid;association_foreignkey:ID"`
	Estado     int         `json:"estado"`
}
