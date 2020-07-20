package structTarea

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Tareapendiente struct {
	structGormModel.GormModel
	Tenant string
}
