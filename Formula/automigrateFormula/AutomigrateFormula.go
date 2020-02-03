package automigrateFormula

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Formula/structFormula"
)

func AutomigrateFormulaTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structFormula.Formula{}).Error
	return err
}
