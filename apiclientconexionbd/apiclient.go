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



type Microservicio interface {
	NecesitaActualizar(*gorm.DB) bool
	AutomigrarPublic(*gorm.DB) error
	AutomigrarPrivate(*gorm.DB) error
	ActualizarVersion(*gorm.DB)
}

var automigratePublicArray = []Microservicio{&automigrateLegajo.MicroservicioLegajo{}, &automigrateConcepto.MicroservicioConcepto{}, &automigrateLiquidacion.MicroservicioLiquidacion{}, &automigrateSiradig.MicroservicioSiradig{}, &automigrateFunction.MicroservicioFunction{}}
var automigratePrivateArray = []Microservicio{&automigrateFunction.MicroservicioFunction{}, &automigrateLegajo.MicroservicioLegajo{}, &automigrateConcepto.MicroservicioConcepto{}, &automigrateNovedad.MicroservicioNovedad{}, &automigrateLiquidacion.MicroservicioLiquidacion{}, &automigrateSiradig.MicroservicioSiradig{}}

func AutomigrateTablaSecurity(db *gorm.DB) (error,bool) {

	actualizo := false
	var err error
	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateAutenticacion.ObtenerVersionAutenticacionConfiguracion(), automigrateAutenticacion.ObtenerVersionAutenticacionDB(db)) {
		if err = automigrateAutenticacion.AutomigrateAutenticacionTablaSecurity(db); err != nil {
			return err, actualizo
		} else {
			actualizo = true
			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateAutenticacion.ObtenerVersionAutenticacionConfiguracion(), automigrateAutenticacion.Security, db)
		}
	}
	return err, actualizo

}

func AutomigrateTablasPublicas(db *gorm.DB) (error, bool) {

	actualizo := false
	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	for _, microservicio := range automigratePublicArray {
		if microservicio.NecesitaActualizar(db) {
			err := microservicio.AutomigrarPublic(db)

			if err != nil {
				return err, false
			} else {
				actualizo = true
				microservicio.ActualizarVersion(db)
			}
		}
	}

	return nil, actualizo
}

func AutomigrateTablasPrivadas(db *gorm.DB) error {

	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	for _, microservicio := range automigratePrivateArray {
		if microservicio.NecesitaActualizar(db) {
			err := microservicio.AutomigrarPrivate(db)

			if err != nil {
				return err
			} else {
				microservicio.ActualizarVersion(db)
			}
		}
	}

	return nil
}

func ForzarActualizacionPrivados() {

}