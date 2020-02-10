package automigrateFunction

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Function/structFunction"
)

func AutomigrateFunctionTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structFunction.Value{}, &structFunction.Invoke{}, &structFunction.Param{}, &structFunction.Function{}).Error
	if err == nil {
		db.Model(&structFunction.Param{}).AddForeignKey("functionname", "function(name)", "CASCADE", "CASCADE")
		db.Model(&structFunction.Function{}).AddForeignKey("valueid", "value(id)", "CASCADE", "CASCADE")
		db.Model(&structFunction.Invoke{}).AddForeignKey("functionname", "function(name)", "CASCADE", "CASCADE")
		db.Model(&structFunction.Value{}).AddForeignKey("valueinvokeid", "invoke(id)", "CASCADE", "CASCADE")

	}
	return err

}
