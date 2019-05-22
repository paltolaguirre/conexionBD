package apiclientconexionbd

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/autenticacion/publico"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

func ObtenerDB(tokenAutenticacion *publico.Security, nombreMicroservicio string, versionMicroservicio int, automigrateTablasPrivadas func(*gorm.DB)) *gorm.DB {

	token := *tokenAutenticacion
	tenant := token.Tenant

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
