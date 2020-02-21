package structFunction

import (
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Param struct {
	structGormModel.GormModel
	Name         string `json:"name"`
	Type         string `json:"type"`
	Functionname string `json:"functionname"`
}
