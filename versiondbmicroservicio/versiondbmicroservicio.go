package versiondbmicroservicio

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Versiondbmicroservicio struct {
	structGormModel.GormModel
	Nombremicroservicio  string
	Versionmicroservicio int
}

func CrearTablaVersionDBMicroservicio(db *gorm.DB) {

	db.AutoMigrate(&Versiondbmicroservicio{})
}

func UltimaVersion(nombre string, db *gorm.DB) int {

	var versiondbmicroservicio Versiondbmicroservicio

	db.First(&versiondbmicroservicio, "nombremicroservicio = ?", nombre)

	return versiondbmicroservicio.Versionmicroservicio

}

func ActualizarMicroservicioVersionDB(db *gorm.DB, versionMicroservicio int, nombreMicroservicio string) {

	var versiondbmicroservicio Versiondbmicroservicio

	db.First(&versiondbmicroservicio, "nombremicroservicio = ?", nombreMicroservicio)

	versiondbmicroservicio.Nombremicroservicio = nombreMicroservicio
	versiondbmicroservicio.Versionmicroservicio = versionMicroservicio

	db.Save(&versiondbmicroservicio)

}

func ActualizarVersionMicroservicio(versionMicroservicioConfiguracion int, versionMicroservicioDB int, nombreMicroservicio string, db *gorm.DB) {
	if versionMicroservicioConfiguracion > versionMicroservicioDB {
		ActualizarMicroservicioVersionDB(db, versionMicroservicioConfiguracion, nombreMicroservicio)
	}
}
