package automigrateLiquidacion

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Liquidacion/structLiquidacion"
)

func AutomigrateLiquidacionTablasPrivadas(db *gorm.DB) error {

	// para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structLiquidacion.Descuento{}, &structLiquidacion.Importenoremunerativo{}, &structLiquidacion.Importeremunerativo{}, &structLiquidacion.Retencion{}, &structLiquidacion.Aportepatronal{}, &structLiquidacion.Liquidacion{}, &structLiquidacion.Liquidacionitem{}).Error
	if err == nil {
		db.Model(&structLiquidacion.Descuento{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Importenoremunerativo{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Importeremunerativo{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Retencion{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Aportepatronal{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Liquidacionitem{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")

		if ObtenerVersionLiquidacionDB(db) < 4 {
			err = unificarDatosEnLaTablaLiquidacionItem(db)
		}

	}
	return err
}

func AutomigrateLiquidacionTablasPublicas(db *gorm.DB) error {
	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structLiquidacion.Liquidacioncondicionpago{}, &structLiquidacion.Liquidaciontipo{}).Error
	return err
}

func unificarDatosEnLaTablaLiquidacionItem(db *gorm.DB) error {
	//abro una transacción para que si hay un error no persista en la DB
	var err error
	tx := db.Begin()
	if err = insertTablaLiquidacionTipo(tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func insertTablaLiquidacionTipo(tx *gorm.DB) error {
	var err error

	if err = tx.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM importeremunerativo)").Error; err != nil {
		return err
	} else {
		tx.Exec("DELETE FROM importeremunerativo")
	}

	if err = tx.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM importenoremunerativo)").Error; err != nil {
		return err
	} else {
		tx.Exec("DELETE FROM importenoremunerativo")

	}

	if err = tx.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM descuento)").Error; err != nil {
		return err
	} else {
		tx.Exec("DELETE FROM descuento")
	}

	if err = tx.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM retencion)").Error; err != nil {
		return err
	} else {
		tx.Exec("DELETE FROM retencion")
	}

	if err = tx.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM aportepatronal)").Error; err != nil {
		return err
	} else {
		tx.Exec("DELETE FROM aportepatronal")
	}

	return err
}
