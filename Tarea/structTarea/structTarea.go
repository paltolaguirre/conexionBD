package structTarea

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Tarea struct {
	structGormModel.GormModel
	Handler       string       `json:"handler"`
	TareaItems    []Tareaitem  `json:"tareaitems" gorm:"ForeignKey:Tareaid;association_foreignkey:ID"`
	Estadotarea   *Estadotarea `json:"estadotarea" gorm:"ForeignKey:estadotareaid;association_foreignkey:ID;association_autoupdate:false;not null;"`
	Estadotareaid *int         `json:"estadotareaid" sql:"type:int REFERENCES EstadoTarea(ID)"`
}
