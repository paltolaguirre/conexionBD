package apiclientconexionbd

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Autenticacion/automigrateAutenticacion"
	"github.com/xubiosueldos/conexionBD/Concepto/automigrateConcepto"
	"github.com/xubiosueldos/conexionBD/Function/automigrateFunction"
	"github.com/xubiosueldos/conexionBD/Legajo/automigrateLegajo"
	"github.com/xubiosueldos/conexionBD/Liquidacion/automigrateLiquidacion"
	"github.com/xubiosueldos/conexionBD/Novedad/automigrateNovedad"
	"github.com/xubiosueldos/conexionBD/Siradig/automigrateSiradig"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
)

func AutomigrateTablaSecurity(db *gorm.DB) error {
	var err error
	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateAutenticacion.ObtenerVersionAutenticacionConfiguracion(), automigrateAutenticacion.ObtenerVersionAutenticacionDB(db)) {
		if err = automigrateAutenticacion.AutomigrateAutenticacionTablaSecurity(db); err != nil {
			return err
		} else {
			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateAutenticacion.ObtenerVersionAutenticacionConfiguracion(), automigrateAutenticacion.Security, db)
		}
	}
	return err

}

func AutomigrateTablasPublicas(db *gorm.DB) error {
	var err error

	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateLegajo.ObtenerVersionLegajoConfiguracion(), automigrateLegajo.ObtenerVersionLegajoDB(db)) {

		if err = automigrateLegajo.AutomigrateLegajoTablasPublicas(db); err != nil {
			return err
		} else {
			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateLegajo.ObtenerVersionLegajoConfiguracion(), automigrateLegajo.Legajo, db)
		}
	}
	if versiondbmicroservicio.ActualizarMicroservicio(automigrateConcepto.ObtenerVersionConceptoConfiguracion(), automigrateConcepto.ObtenerVersionConceptoDB(db)) {

		if err = automigrateConcepto.AutomigrateConceptoTablasPublicas(db); err != nil {
			return err
		} else {
			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateConcepto.ObtenerVersionConceptoConfiguracion(), automigrateConcepto.Concepto, db)
		}
	}

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateLiquidacion.ObtenerVersionLiquidacionConfiguracion(), automigrateLiquidacion.ObtenerVersionLiquidacionDB(db)) {

		if err = automigrateLiquidacion.AutomigrateLiquidacionTablasPublicas(db); err != nil {
			return err
		} else {

			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateLiquidacion.ObtenerVersionLiquidacionConfiguracion(), automigrateLiquidacion.Liquidacion, db)
		}
	}

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateSiradig.ObtenerVersionSiradigConfiguracion(), automigrateSiradig.ObtenerVersionSiradigDB(db)) {

		if err = automigrateSiradig.AutomigrateSiradigTablasPublicas(db); err != nil {
			return err
		} else {

			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateSiradig.ObtenerVersionSiradigConfiguracion(), automigrateSiradig.Siradig, db)
		}
	}

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateFunction.ObtenerVersionFunctionConfiguracion(), automigrateFunction.ObtenerVersionFunctionDB(db)) {

		if err = automigrateFunction.AutomigrateFunctionTablasPublicas(db); err != nil {
			return err
		} else {

			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateFunction.ObtenerVersionFunctionConfiguracion(), automigrateFunction.Function, db)
		}
	}

	return err
}

func AutomigrateTablasPrivadas(db *gorm.DB) error {
	var err error

	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateFunction.ObtenerVersionFunctionConfiguracion(), automigrateFunction.ObtenerVersionFunctionDB(db)) {
		if err = automigrateFunction.AutomigrateFunctionTablasPrivadas(db); err != nil {
			return err
		} else {
			if err = automigrateFunction.ObtenerFormulasPublicas(db); err != nil {
				return err
			}
			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateFunction.ObtenerVersionFunctionConfiguracion(), automigrateFunction.Function, db)
		}
	}

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateLegajo.ObtenerVersionLegajoConfiguracion(), automigrateLegajo.ObtenerVersionLegajoDB(db)) {

		if err = automigrateLegajo.AutomigrateLegajoTablasPrivadas(db); err != nil {
			return err
		} else {
			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateLegajo.ObtenerVersionLegajoConfiguracion(), automigrateLegajo.Legajo, db)
		}
	}

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateConcepto.ObtenerVersionConceptoConfiguracion(), automigrateConcepto.ObtenerVersionConceptoDB(db)) {

		if err = automigrateConcepto.AutomigrateConceptoTablasPrivadas(db); err != nil {
			return err
		} else {
			if err = automigrateConcepto.ObtenerConceptosPublicos(db); err != nil {
				return err
			}
			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateConcepto.ObtenerVersionConceptoConfiguracion(), automigrateConcepto.Concepto, db)
		}
	}

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateNovedad.ObtenerVersionNovedadConfiguracion(), automigrateNovedad.ObtenerVersionNovedadDB(db)) {
		if err = automigrateNovedad.AutomigrateNovedadTablasPrivadas(db); err != nil {
			return err
		} else {
			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateNovedad.ObtenerVersionNovedadConfiguracion(), automigrateNovedad.Novedad, db)
		}
	}

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateLiquidacion.ObtenerVersionLiquidacionConfiguracion(), automigrateLiquidacion.ObtenerVersionLiquidacionDB(db)) {
		if err = automigrateLiquidacion.AutomigrateLiquidacionTablasPrivadas(db); err != nil {
			return err
		} else {

			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateLiquidacion.ObtenerVersionLiquidacionConfiguracion(), automigrateLiquidacion.Liquidacion, db)
		}
	}

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateSiradig.ObtenerVersionSiradigConfiguracion(), automigrateSiradig.ObtenerVersionSiradigDB(db)) {
		if err = automigrateSiradig.AutomigrateSiradigTablasPrivadas(db); err != nil {
			return err
		} else {

			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateSiradig.ObtenerVersionSiradigConfiguracion(), automigrateSiradig.Siradig, db)
		}
	}

	return err
}
