package structFunction

import (
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Value struct {
	structGormModel.GormModel
	Name          string  `json:"name"`
	Valuenumber   float64 `json:"valuenumber" sql:"type:decimal(19,4);"`
	Valuestring   string  `json:"valuestring"`
	Valueboolean  bool    `json:"Valueboolean"`
	Valueinvoke   *Invoke `json:"valueinvoke" gorm:"ForeignKey:Valueinvokeid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Valueinvokeid int     `json:"valueinvokeid"`
	Arginvokeid   int     `json:"arginvokeid"`
	// Valueobject   interface{} `json:"valueobject" gorm:"type:json"`
}
