package automigrateConcepto

import (
	"net/http"

	"github.com/xubiosueldos/conexionBD"

	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Concepto/structConcepto"
	"github.com/xubiosueldos/framework"
)

func AutomigrateConceptoTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structConcepto.Concepto{}).Error
	return err
}

func AutomigrateConceptoTablasPublicas(db *gorm.DB) error {
	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structConcepto.Tipoconcepto{}, &structConcepto.Tipodecalculo{}, &structConcepto.Tipoimpuestoganancias{}, &structConcepto.Concepto{}).Error
	if err == nil {

		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES (current_timestamp,'Remuneración bruta','REMUNERACION_BRUTA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Retribuciones no habituales','RETRIBUCIONES_NO_HABITUALES','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Horas extras remuneración gravada','HORAS_EXTRAS_REMUNERACION_GRAVADA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Movilidad y viáticos remuneración gravada','MOVILIDAD_Y_VIATICOS_REMUNERACION_GRAVADA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Material didáctico personal docente remuneración gravada','MATERIAL_DIDACTICO_PERSONAL_DOCENTE_REMUNERACION_GRAVADA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Remuneración no alcanzada o exenta','REMUNERACION_NO_ALCANZADA_O_EXENTA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Horas extras remuneración exenta','HORAS_EXTRAS_REMUNERACION_EXENTA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Movilidad y viáticos remuneración exenta','MOVILIDAD_Y_VIATICOS_REMUNERACION_EXENTA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Material didáctico personal docente remuneración exenta','MATERIAL_DIDACTICO_PERSONAL_DOCENTE_REMUNERACION_EXENTA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Aportes jubilatorios, retiros, pensiones o subsidios','APORTES_JUBILATORIOS_RETIROS_PENSIONES_O_SUBSIDIOS','',1, false, false, false, true, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Aportes obra social','APORTES_OBRA_SOCIAL','',1, false, false, false, true, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Cuota sindical','CUOTA_SINDICAL','',1, false, false, false, true, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Descuentos obligatorios por ley nacional, provincial o municipal','DESCUENTOS_OBLIGATORIOS_POR_LEY_NACIONAL_PROVINCIAL_MUNICIPAL','',1, false, false, false, true, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Gastos movilidad, viáticos abonados por el empleador','GASTOS_MOVILIDAD_VIATICOS_ABONADOS_POR_EL_EMPLEADOR','',1, false, false, false, true, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(current_timestamp,'Otras deducciones','OTRAS_DEDUCCIONES','',1, false, false, false, true, false)")

		db.Exec("UPDATE public.Tipoimpuestoganancias SET id = -id;")

		db.Exec("UPDATE CONCEPTO SET prorateo = false, basesac = true, tipoimpuestogananciasid = -1 WHERE ID IN (-1,-3,-4,-15,-16,-17)")
		db.Exec("UPDATE CONCEPTO SET prorateo = false, basesac = true, tipoimpuestogananciasid = -3 WHERE ID IN (-5,-6)")
		db.Exec("UPDATE CONCEPTO SET prorateo = false, basesac = false, tipoimpuestogananciasid = -2 WHERE ID IN (-8,-9,-10,-11,-13,-14)")
		db.Exec("UPDATE CONCEPTO SET prorateo = true, basesac = true, tipoimpuestogananciasid = -2 WHERE ID = -12")
		db.Exec("UPDATE CONCEPTO SET prorateo = false, basesac = false, tipoimpuestogananciasid = -10 WHERE ID IN (-18,-19)")
		db.Exec("UPDATE CONCEPTO SET prorateo = false, basesac = false, tipoimpuestogananciasid = -11 WHERE ID = -20")

	}
	return err
}

func ObtenerConceptosPublicos(db *gorm.DB) {
	var w http.ResponseWriter
	var conceptos []structConcepto.Concepto

	// nil hace referencia a la funcion automigrate
	db_public := conexionBD.ObtenerDB("public")

	db_public.Find(&conceptos)
	for i := 0; i < len(conceptos); i++ {
		concepto := conceptos[i]
		if err := db.Save(&concepto).Error; err != nil {
			framework.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

}
