package structTarea

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Cancelacion struct {
	structGormModel.GormModel
	Tarea   *Tarea `json:"tarea" gorm:"ForeignKey:tareaid;association_foreignkey:ID;association_autoupdate:false;not null;"`
	Tareaid *int  `json:"tareaid" sql:"unique;"`
}
