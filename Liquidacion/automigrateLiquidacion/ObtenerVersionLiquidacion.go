package automigrateLiquidacion

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

func ObtenerVersionConceptoDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion("liquidacion", db)
}

func ObtenerVersionLiquidacionConfiguracion(db *gorm.DB) int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionliquidacion
}
