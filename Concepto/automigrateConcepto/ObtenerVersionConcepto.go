package automigrateConcepto

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

const (
	Concepto = "concepto"
)

func ObtenerVersionConceptoDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(Concepto, db)
}

func ObtenerVersionConceptoConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionconcepto
}
