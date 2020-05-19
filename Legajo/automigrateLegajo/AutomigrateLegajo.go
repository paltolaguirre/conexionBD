package automigrateLegajo

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
)

type MicroservicioLegajo struct{
}

func (*MicroservicioLegajo) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(ObtenerVersionLegajoConfiguracion(), ObtenerVersionLegajoDB(db))
}

func (*MicroservicioLegajo) AutomigrarPublic(db *gorm.DB) error {
	return AutomigrateLegajoTablasPublicas(db)
}

func (*MicroservicioLegajo) AutomigrarPrivate(db *gorm.DB) error {
	return AutomigrateLegajoTablasPrivadas(db)
}

func (*MicroservicioLegajo) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(ObtenerVersionLegajoConfiguracion(), Legajo, db)
}

func AutomigrateLegajoTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structLegajo.Conyuge{}, &structLegajo.Hijo{}, &structLegajo.Legajo{}).Error
	if err == nil {
		db.Model(&structLegajo.Hijo{}).AddForeignKey("legajoid", "legajo(id)", "CASCADE", "CASCADE")
		db.Model(&structLegajo.Conyuge{}).AddForeignKey("legajoid", "legajo(id)", "CASCADE", "CASCADE")
	}
	return err
}

func AutomigrateLegajoTablasPublicas(db *gorm.DB) error {
	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structLegajo.Pais{}, &structLegajo.Provincia{}, &structLegajo.Localidad{}, &structLegajo.Modalidadcontratacion{}, &structLegajo.Situacion{}, &structLegajo.Condicion{}, &structLegajo.Condicionsiniestrado{}, &structLegajo.Obrasocial{}, &structLegajo.Estadocivil{}).Error
	return err
}
