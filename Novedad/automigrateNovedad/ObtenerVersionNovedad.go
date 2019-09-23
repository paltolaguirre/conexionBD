package automigrateNovedad

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

const (
	Novedad = "novedad"
)

func ObtenerVersionNovedadDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(Novedad, db)
}

func ObtenerVersionNovedadConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionnovedad
}
