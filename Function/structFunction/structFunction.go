package structFunction

import (
	"time"
)

type Function struct {
	Name        string `json:"name" gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	Params      []Param    `json:"params" gorm:"ForeignKey:Functionname;association_foreignkey:Name"`
	Description string     `json:"description"`
	Origin      string     `json:"origin"`
	Type        string     `json:"type"`
	Scope       string     `json:"scope"`
	Result      string     `json:"result"`
	Value       *Value     `json:"value" gorm:"ForeignKey:Valueid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Valueid     *int       `json:"valueid" gorm:"not null"`
}
