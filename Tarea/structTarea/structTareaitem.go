package structTarea

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Tareaitem struct {
	structGormModel.GormModel
	Tareaid           int              `json:"tareaid"`
	Json              string           `json:"json" sql:"type:JSON;"`
	Estadotareaitem   *Estadotareaitem `json:"estadotareaitem" gorm:"ForeignKey:estadotareaitemid;association_foreignkey:ID;association_autoupdate:false;not null;"`
	Estadotareaitemid *int             `json:"estadotareaitemid" sql:"type:int REFERENCES Estadotareaitem(ID)"`
	Detalle           string           `json:"detalle"`
}
