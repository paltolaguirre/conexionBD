package automigrateLiquidacion

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Liquidacion/structLiquidacion"
)

func AutomigrateLiquidacionTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structLiquidacion.Descuento{}, &structLiquidacion.Importenoremunerativo{}, &structLiquidacion.Importeremunerativo{}, &structLiquidacion.Retencion{}, &structLiquidacion.Aportepatronal{}, &structLiquidacion.Liquidacion{}, &structLiquidacion.Liquidacionitem{}).Error
	if err == nil {
		db.Model(&structLiquidacion.Descuento{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Importenoremunerativo{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Importeremunerativo{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Retencion{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Aportepatronal{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Liquidacionitem{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")

		//unificarDatosEnLaTablaLiquidacionItem(db)
	}
	return err
}

func AutomigrateLiquidacionTablasPublicas(db *gorm.DB) error {
	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structLiquidacion.Liquidacioncondicionpago{}, &structLiquidacion.Liquidaciontipo{}).Error
	return err
}

/*func unificarDatosEnLaTablaLiquidacionItem(db *gorm.DB) {
	//Se podria agregar otra columna en cada una de las tablas a unificar para saber si el registro fue o no colocado en la tabla liquidacionitem
	db.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM importeremunerativo)")
	db.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM importenoremunerativo)")
	db.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM descuento)")
	db.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM retencion)")
	db.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM aportepatronal)")
}*/
