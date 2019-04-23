package conexionBD

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xubiosueldos/legajo/structLegajo"
)

var db *gorm.DB
var err error

func ConnectBD(tenant string) *gorm.DB {

	db, err = gorm.Open("postgres", "host=192.168.30.111 port=5432 user=postgres dbname=faf_multitenant_go password=Post66MM/")

	if err != nil {
		panic("failed to connect database")
	}
	db.Exec("SET search_path = " + tenant + ",public")

	db.SingularTable(true)

	db.AutoMigrate(&structLegajo.Pais{}, &structLegajo.Provincia{}, &structLegajo.Localidad{}, &structLegajo.Zona{}, &structLegajo.Modalidadcontratacion{}, &structLegajo.Situacion{}, &structLegajo.Condicion{}, &structLegajo.Condicionsiniestrado{}, &structLegajo.Conveniocolectivo{}, &structLegajo.Centrodecosto{}, &structLegajo.Obrasocial{}, &structLegajo.Conyuge{}, &structLegajo.Hijo{}, &structLegajo.Legajo{})

	db.Model(&structLegajo.Hijo{}).AddForeignKey("legajoid", "legajo(id)", "CASCADE", "CASCADE")
	db.Model(&structLegajo.Conyuge{}).AddForeignKey("legajoid", "legajo(id)", "CASCADE", "CASCADE")

	return db
}
