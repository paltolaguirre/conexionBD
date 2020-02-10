package structFunction

import (
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Value struct {
	structGormModel.GormModel
	Name          string  `json:"name"`
	Valuenumber   int64   `json:"valuenumber"`
	Valuestring   string  `json:"valuestring"`
	Valueinvoke   *Invoke `json:"valueinvoke" gorm:"ForeignKey:Valueinvokeid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Valueinvokeid int     `json:"valueinvokeid"`
	Arginvokeid   int     `json:"arginvokeid"`
}
