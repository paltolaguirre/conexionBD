package automigrateLiquidacion

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

const (
	Liquidacion = "liquidacion"
)

func ObtenerVersionLiquidacionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(Liquidacion, db)
}

func ObtenerVersionLiquidacionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionliquidacion
}
