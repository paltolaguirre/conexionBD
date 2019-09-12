package automigrateLegajo

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

func ObtenerVersionLegajoDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion("legajo", db)
}

func ObtenerVersionLegajoConfiguracion(db *gorm.DB) int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionlegajo
}
