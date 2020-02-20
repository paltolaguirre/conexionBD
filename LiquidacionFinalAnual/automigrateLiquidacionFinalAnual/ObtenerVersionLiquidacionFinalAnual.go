package automigrateLiquidacionFinalAnual

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

const (
	Liquidacionfinalanual = "liquidacionfinalanual"
)

func ObtenerVersionLiquidacionfinalanualDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(Liquidacionfinalanual, db)
}

func ObtenerVersionLiquidacionfinalanualConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionliquidacionfinalanual
}
