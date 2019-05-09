package conexionBD

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xubiosueldos/framework/configuracion"
)

var db *gorm.DB
var err error

func ConnectBD(tenant string) *gorm.DB {

	configuracion := configuracion.GetInstance()

	db, err = gorm.Open("postgres", "host= "+configuracion.Ip+"port=5432 user=postgres dbname="+configuracion.Namedb+"password=Post66MM/")

	if err != nil {
		panic("failed to connect database")
	}

	//Crea el schema si no existe
	db.Exec("CREATE SCHEMA IF NOT EXISTS " + tenant)

	db.SingularTable(true)

	if tenant == "public" {

		db.Exec("SET search_path = " + tenant)

	} else {
		db.Exec("SET search_path = " + tenant + ",public")
	}

	return db
}
