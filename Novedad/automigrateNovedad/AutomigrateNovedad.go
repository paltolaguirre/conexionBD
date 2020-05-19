package automigrateNovedad

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Novedad/structNovedad"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
)

type MicroservicioNovedad struct{
}

func (*MicroservicioNovedad) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(ObtenerVersionNovedadConfiguracion(), ObtenerVersionNovedadDB(db))
}

func (*MicroservicioNovedad) AutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*MicroservicioNovedad) AutomigrarPrivate(db *gorm.DB) error {
	return AutomigrateNovedadTablasPrivadas(db)
}

func (*MicroservicioNovedad) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(ObtenerVersionNovedadConfiguracion(), Novedad, db)
}

func AutomigrateNovedadTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structNovedad.Novedad{}).Error
	db.Model(&structNovedad.Novedad{}).AddForeignKey("conceptoid", "concepto(id)", "RESTRICT", "RESTRICT")
	if ObtenerVersionNovedadDB(db) < 2 {
		// err = db.Exec("alter table novedad alter column cantidad type numeric(19,4);").Error
		db.Exec("alter table novedad alter column cantidad type numeric(19,4);")
	}
	return err
}
