package automigrateAutenticacion

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

func ObtenerVersionAutenticacionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion("security", db)
}

func ObtenerVersionAutenticacionConfiguracion(db *gorm.DB) int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionsecurity
}
