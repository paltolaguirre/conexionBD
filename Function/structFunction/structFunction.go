package structFunction

import (
	"time"

	"github.com/jinzhu/gorm"
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
	Value       *Value     `json:"value" gorm:"ForeignKey:Valueid;association_foreignkey:ID;association_autoupdate:false;"`
	Valueid     *int       `json:"valueid" sql:"type:int REFERENCES Value(ID)"`
}

// AfterDelete hook defined for cascade delete
func (function *Function) AfterDelete(tx *gorm.DB) error {
	return tx.Model(&Value{}).Where("id = ?", function.Valueid).Unscoped().Delete(&Value{}).Error
}
