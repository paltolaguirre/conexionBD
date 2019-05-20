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

func ActualizarVersionMicroservicio(db *gorm.DB, versionMicroservicio int, nombreMicroservicio string) {

	var versiondbmicroservicio Versiondbmicroservicio

	db.First(&versiondbmicroservicio, "nombremicroservicio = ?", nombreMicroservicio)

	versiondbmicroservicio.Nombremicroservicio = nombreMicroservicio
	versiondbmicroservicio.Versionmicroservicio = versionMicroservicio

	db.Save(&versiondbmicroservicio)

}
