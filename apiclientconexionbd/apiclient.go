package apiclientconexionbd

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

func ObtenerDB(tenant string, nombreMicroservicio string, versionMicroservicio int, automigrateTablasPrivadas func(*gorm.DB)) *gorm.DB {

	db := conexionBD.ConnectBD(tenant)

	crearTablaVersionMicroServicioYPrivadas(nombreMicroservicio, versionMicroservicio, automigrateTablasPrivadas, db)

	return db
}
func crearTablaVersionMicroServicioYPrivadas(nombreMicroservicio string, versionMicroservicio int, automigrateTablasPrivadas func(*gorm.DB), db *gorm.DB) {

	configuracion := configuracion.GetInstance()

	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	if versionMicroservicio > versiondbmicroservicio.UltimaVersion(nombreMicroservicio, db) {
		automigrateTablasPrivadas(db)
		versiondbmicroservicio.ActualizarVersionMicroservicio(db, configuracion.Versionlegajo, nombreMicroservicio)
	}

}

func CerrarDB(db *gorm.DB) {
	db.DB().Close()
}
