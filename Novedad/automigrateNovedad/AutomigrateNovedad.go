package automigrateNovedad

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Novedad/structNovedad"
)

func AutomigrateNovedadTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structNovedad.Novedad{}).Error
	if ObtenerVersionNovedadDB(db) < 2 {
		err = db.Exec("alter table novedad alter column cantidad type numeric(19,4);").Error
	}
	return err
}
