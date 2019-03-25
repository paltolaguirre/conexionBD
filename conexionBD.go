package conexionBD

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func ConnectBD(tenant string) *gorm.DB {

	db, err = gorm.Open("postgres", "host=192.168.30.111 port=5432 user=postgres dbname=DES2_MULTITENANT_AR_1 password=Post66MM/")

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
