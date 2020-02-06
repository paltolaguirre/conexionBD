package structFunction

import (
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Invoke struct {
	structGormModel.GormModel
	Function     *Function `json:"function" gorm:"ForeignKey:Functionname;association_foreignkey:name;association_autoupdate:false;not null"`
	Functionname string    `json:"functionname"`
	Args         []Value   `json:"args" gorm:"ForeignKey:Invokeid;association_foreignkey:ID"`
}
