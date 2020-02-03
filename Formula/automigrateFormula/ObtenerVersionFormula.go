package automigrateFormula

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

const (
	Formula = "formula"
)

func ObtenerVersionFormulaDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(Formula, db)
}

func ObtenerVersionFormulaConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionformula
}
