package automigrateLiquidacionFinalAnual

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/LiquidacionFinalAnual/structLiquidacionFinalAnual"
)

func AutomigrateLiquidacionFinalAnualTablasPrivadas(db *gorm.DB) error {
	err := db.AutoMigrate(&structLiquidacionFinalAnual.Liquidacionfinalanual{}).Error
	return err
}

func AutomigrateLiquidacionFinalAnualTablasPublicas(db *gorm.DB) error {
	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structLiquidacionFinalAnual.Tipopresentacionliquidacionfinalanual{}).Error
	if err == nil {

		db.Exec("INSERT INTO TIPOPRESENTACIONLIQUIDACIONFINALANUAL(created_at, nombre, codigo, descripcion, activo) VALUES(current_timestamp,'Anual','ANUAL','',1)")
		db.Exec("INSERT INTO TIPOPRESENTACIONLIQUIDACIONFINALANUAL(created_at, nombre, codigo, descripcion, activo) VALUES(current_timestamp,'Final','FINAL','',1)")
		db.Exec("UPDATE TIPOPRESENTACIONLIQUIDACIONFINALANUAL SET id = -id ")
	}
	return err
}
