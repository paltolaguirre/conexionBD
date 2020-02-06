package structFunction

import (
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Value struct {
	structGormModel.GormModel
	Name        string  `json:"name"`
	Valuenumber int64   `json:"valuenumber"`
	Valuestring string  `json:"valuestring"`
	Invoke      *Invoke `json:"invoke" gorm:"ForeignKey:Invokeid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Invokeid    int     `json:"invokeid"`
}
