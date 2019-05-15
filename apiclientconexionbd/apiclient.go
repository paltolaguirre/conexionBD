package apiclientconexionbd

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/autenticacion/publico"
	"github.com/xubiosueldos/conexionBD"
)

func obtenerDB(tokenAutenticacion *publico.Security) *gorm.DB {

	token := *tokenAutenticacion
	tenant := token.Tenant

	return conexionBD.ConnectBD(tenant)

}
