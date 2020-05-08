package automigrateFunction

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

const (
	Function = "formula"
)

func ObtenerVersionFunctionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(Function, db)
}

func ObtenerVersionFunctionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionformula
}
