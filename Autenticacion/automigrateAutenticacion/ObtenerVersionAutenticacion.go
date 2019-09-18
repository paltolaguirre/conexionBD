package automigrateAutenticacion

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

const (
	Security = "security"
)

func ObtenerVersionAutenticacionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(Security, db)
}

func ObtenerVersionAutenticacionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionsecurity
}
