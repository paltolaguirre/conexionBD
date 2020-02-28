package conexionBD

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xubiosueldos/framework/configuracion"
)

func ConnectBD(tenant string) *gorm.DB {
	var db *gorm.DB
	var err error

	configuracion := configuracion.GetInstance()

	db, err = gorm.Open("postgres", "host= "+configuracion.Ip+" port=5432 user=postgres dbname= "+configuracion.Namedb+" password="+configuracion.Passdb+" sslmode="+configuracion.Ssldb)

	if err != nil {
		panic("failed to connect database")
	}
	db.DB().SetConnMaxLifetime(time.Second * 120)
	//db.DB().SetMaxIdleConns()
	//db.DB().SetMaxOpenConns()
	//Crea el schema si no existe
	db.Exec("CREATE SCHEMA IF NOT EXISTS " + tenant)
	db.Exec("GRANT USAGE ON SCHEMA " + tenant + " TO read_only")
	db.Exec("ALTER DEFAULT PRIVILEGES FOR USER postgres	IN SCHEMA " + tenant +" GRANT SELECT ON TABLES TO read_only")


	db.SingularTable(true)

	if tenant == "public" {

		db.Exec("SET search_path = " + tenant)

	} else {
		db.Exec("SET search_path = " + tenant + ",public")
	}

	return db
}

func ObtenerDB(tenant string) *gorm.DB {
	return ConnectBD(tenant)
}

func CerrarDB(db *gorm.DB) {
	db.DB().Close()
}
