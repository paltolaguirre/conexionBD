package automigrateAutenticacion

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Autenticacion/structAutenticacion"
)

func AutomigrateAutenticacionTablaSecurity(db *gorm.DB) error {
	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structAutenticacion.Security{}).Error
	return err
}
