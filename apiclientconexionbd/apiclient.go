package apiclientconexionbd

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Autenticacion/automigrateAutenticacion"
	"github.com/xubiosueldos/conexionBD/Concepto/automigrateConcepto"
	"github.com/xubiosueldos/conexionBD/Legajo/automigrateLegajo"
	"github.com/xubiosueldos/conexionBD/Liquidacion/automigrateLiquidacion"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
)

func AutomigrateTablaSecurity(db *gorm.DB) {
	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)
	err := automigrateAutenticacion.AutomigrateAutenticacionTablasPublicas(db)
	if err == nil {
		ActualizarVersionMicroservicio(automigrateAutenticacion.ObtenerVersionAutenticacionConfiguracion(), automigrateAutenticacion.ObtenerVersionAutenticacionDB(), "security", db)
	}

}

func AutomigrateTablasPublicas(db *gorm.DB) {
	var err error
	var versionMicroservicio int

	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	err = automigrateLegajo.AutomigrateLegajoPublicas(db)
	if err == nil {
		ActualizarVersionMicroservicio(automigrateLegajo.ObtenerVersionLegajoConfiguracion(), automigrateLegajo.ObtenerVersionLegajoDB(), "legajo", db)
	}

	err = automigrateConcepto.AutomigrateConceptoPublicas(db)
	if err == nil {
		ActualizarVersionMicroservicio(automigrateConcepto.ObtenerVersionConceptoConfiguracion(), automigrateConcepto.ObtenerVersionConceptoDB(), "concepto", db)
	}

	err = automigrateLiquidacion.AutomigrateLiquidacionPublicas(db)
	if err == nil {

		ActualizarVersionMicroservicio(automigrateLiquidacion.ObtenerVersionLiquidacionConfiguracion(), automigrateLiquidacion.ObtenerVersionLiquidacionDB(), "liquidacion", db)
	}
}

func AutomigrateTablasPrivadas(db *gorm.DB) {
	var err error
	var versionMicroservicio int

	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	err = automigrateLegajo.AutomigrateLegajoPrivadas(db)
	if err == nil {
		ActualizarVersionMicroservicio(automigrateLegajo.ObtenerVersionLegajoConfiguracion(), automigrateLegajo.ObtenerVersionLegajoDB(), "legajo", db)
	}

	err = automigrateConcepto.AutomigrateConceptoPrivadas(db)
	if err == nil {
		ActualizarVersionMicroservicio(automigrateConcepto.ObtenerVersionConceptoConfiguracion(), automigrateConcepto.ObtenerVersionConceptoDB(), "concepto", db)
	}

	err = automigrateNovedad.AutomigrateNovedadPrivadas(db)
	if err == nil {
		ActualizarVersionMicroservicio(automigrateNovedad.ObtenerVersionConceptoConfiguracion(), automigrateNovedad.ObtenerVersionConceptoDB(), "novedad", db)
	}

	err = automigrateLiquidacion.AutomigrateLiquidacionPrivadas(db)
	if err == nil {

		ActualizarVersionMicroservicio(automigrateLiquidacion.ObtenerVersionLiquidacionConfiguracion(), automigrateLiquidacion.ObtenerVersionLiquidacionDB(), "liquidacion", db)
	}

}

func ObtenerDB(tenant string) *gorm.DB {
	return conexionBD.ConnectBD(tenant)
}

func CerrarDB(db *gorm.DB) {
	db.DB().Close()
}
