package automigrateConcepto

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Concepto/structConcepto"
)

func AutomigrateConceptoTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structConcepto.Concepto{}).Error
	if err == nil {
		db.Model(&structConcepto.Concepto{}).AddForeignKey("tipoconceptoid", "tipoconcepto(id)", "RESTRICT", "RESTRICT")
		db.Model(&structConcepto.Concepto{}).AddForeignKey("tipodecalculoid", "tipodecalculo(id)", "RESTRICT", "RESTRICT")
		db.Model(&structConcepto.Concepto{}).AddForeignKey("tipoimpuestogananciasid", "tipoimpuestoganancias(id)", "RESTRICT", "RESTRICT")
		db.Model(&structConcepto.Concepto{}).AddForeignKey("tipocalculoautomaticoid", "tipocalculoautomatico(id)", "RESTRICT", "RESTRICT")
		db.Model(&structConcepto.Concepto{}).AddForeignKey("formulanombre", "function(name)", "RESTRICT", "RESTRICT")

		versionConceptoDB := ObtenerVersionConceptoDB(db)

		if versionConceptoDB < 10 {
			db.Exec("update concepto set tipocalculoautomaticoid = -1 where tipodecalculoid is null")
			db.Exec("update concepto set tipocalculoautomaticoid = -2 where tipodecalculoid is not null")
			db.Exec("update concepto set eseditable = false where tipocalculoautomaticoid != -1")
			db.Exec("update concepto set eseditable = true where tipocalculoautomaticoid = -1")
			/*
				db.Exec("update concepto set formulanombre = 'ImpuestoALasGanancias', tipocalculoautomaticoid = -3 where id = -29")
				db.Exec("update concepto set formulanombre = 'ImpuestoALasGananciasDevolucion', tipocalculoautomaticoid = -3 where id = -30")

				db.Exec("UPDATE CONCEPTO SET eseditable = false WHERE id in (-29, -30)")
			*/
		}

		if versionConceptoDB < 13 {
			db.Exec("update concepto set esremvariable = false where esremvariable is NULL")
		}
	}

	return err
}

func AutomigrateConceptoTablasPublicas(db *gorm.DB) error {
	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structConcepto.Tipocalculoautomatico{}, &structConcepto.Tipoconcepto{}, &structConcepto.Tipodecalculo{}, &structConcepto.Tipoimpuestoganancias{}, &structConcepto.Concepto{}).Error
	if err == nil {

		versionConceptoDB := ObtenerVersionConceptoDB(db)

		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-16,current_timestamp,'SAC','SAC','',1, false, false, false, true, false)")
		db.Exec("UPDATE CONCEPTO SET tipoimpuestogananciasid = -16 WHERE id = -2")

		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-1,current_timestamp,'Remuneración bruta','REMUNERACION_BRUTA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-2,current_timestamp,'Retribuciones no habituales','RETRIBUCIONES_NO_HABITUALES','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-3,current_timestamp,'Horas extras remuneración gravada','HORAS_EXTRAS_REMUNERACION_GRAVADA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-4,current_timestamp,'Movilidad y viáticos remuneración gravada','MOVILIDAD_Y_VIATICOS_REMUNERACION_GRAVADA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-5,current_timestamp,'Material didáctico personal docente remuneración gravada','MATERIAL_DIDACTICO_PERSONAL_DOCENTE_REMUNERACION_GRAVADA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-6,current_timestamp,'Remuneración no alcanzada o exenta','REMUNERACION_NO_ALCANZADA_O_EXENTA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-7,current_timestamp,'Horas extras remuneración exenta','HORAS_EXTRAS_REMUNERACION_EXENTA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-8,current_timestamp,'Movilidad y viáticos remuneración exenta','MOVILIDAD_Y_VIATICOS_REMUNERACION_EXENTA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-9,current_timestamp,'Material didáctico personal docente remuneración exenta','MATERIAL_DIDACTICO_PERSONAL_DOCENTE_REMUNERACION_EXENTA','',1, true, true, true, false, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-10,current_timestamp,'Aportes jubilatorios, retiros, pensiones o subsidios','APORTES_JUBILATORIOS_RETIROS_PENSIONES_O_SUBSIDIOS','',1, false, false, false, true, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-11,current_timestamp,'Aportes obra social','APORTES_OBRA_SOCIAL','',1, false, false, false, true, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-12,current_timestamp,'Cuota sindical','CUOTA_SINDICAL','',1, false, false, false, true, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-13,current_timestamp,'Descuentos obligatorios por ley nacional, provincial o municipal','DESCUENTOS_OBLIGATORIOS_POR_LEY_NACIONAL_PROVINCIAL_MUNICIPAL','',1, false, false, false, true, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-14,current_timestamp,'Gastos movilidad, viáticos abonados por el empleador','GASTOS_MOVILIDAD_VIATICOS_ABONADOS_POR_EL_EMPLEADOR','',1, false, false, false, true, false)")
		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-15,current_timestamp,'Otras deducciones','OTRAS_DEDUCCIONES','',1, false, false, false, true, false)")

		db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid) VALUES(-29, current_timestamp,'Impuesto a las Ganancias', 'IMPUESTO_GANANCIAS',  '', 1, '',-46, true, -4,false, null, null, false, false, -1)")
		db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid) VALUES(-30, current_timestamp,'Impuesto a las Ganancias (Devolución)', 'IMPUESTO_GANANCIAS_DEVOLUCION',  '', 1, '',-46, true, -2,false, null, null, false, false, -1)")
		db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid) VALUES(-31, current_timestamp,'Cuota Sindical', 'CUOTA_SINDICAL',  '', 1, '',-46, true, -4,false, 2, -3, false, true, -12)")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -18")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -19")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -20")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -31")
		//db.Exec("UPDATE CONCEPTO SET eseditable = false WHERE id in (-29, -30)")
		db.Exec("INSERT INTO TIPOCALCULOAUTOMATICO(id, created_at, nombre, codigo, descripcion, activo) VALUES(-1, current_timestamp, 'No Aplica', 'NO_APLICA', '', 1)")
		db.Exec("INSERT INTO TIPOCALCULOAUTOMATICO(id, created_at, nombre, codigo, descripcion, activo) VALUES(-2, current_timestamp, 'Porcentaje', 'PORCENTAJE', '', 1)")
		db.Exec("INSERT INTO TIPOCALCULOAUTOMATICO(id, created_at, nombre, codigo, descripcion, activo) VALUES(-3, current_timestamp, 'Formula', 'FORMULA', '', 1)")
		db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid) VALUES(-32, current_timestamp,'Incremento Salarial Dto 14/20', 'INCREMENTO_SALARIAL_DTO_14_20',  '', 1, '',-46, true, -1,false, null, null, false, true, -1)")

		if versionConceptoDB < 10 {
			err = db.Exec("update concepto set tipocalculoautomaticoid = -1 where tipodecalculoid is null").Error
			err = db.Exec("update concepto set tipocalculoautomaticoid = -2 where tipodecalculoid is not null").Error

			err = db.Exec("update concepto set formulanombre = 'ImpuestoALasGanancias', tipocalculoautomaticoid = -3 where id = -29").Error
			err = db.Exec("update concepto set formulanombre = 'ImpuestoALasGananciasDevolucion', tipocalculoautomaticoid = -3 where id = -30").Error

			err = db.Exec("update concepto set eseditable = false where tipocalculoautomaticoid != -1").Error
			err = db.Exec("update concepto set eseditable = true where tipocalculoautomaticoid = -1").Error
		}

		db.Exec("UPDATE CONCEPTO SET prorrateo = false, basesac = true, tipoimpuestogananciasid = -1 WHERE ID IN (-1,-3,-4,-15,-16,-17)")
		db.Exec("UPDATE CONCEPTO SET tipoimpuestogananciasid = -1 WHERE ID = -2")
		db.Exec("UPDATE CONCEPTO SET prorrateo = false, basesac = true, tipoimpuestogananciasid = -3 WHERE ID IN (-5,-6)")
		db.Exec("UPDATE CONCEPTO SET prorrateo = false, basesac = false, tipoimpuestogananciasid = -2 WHERE ID IN (-8,-9,-10,-11,-13,-14)")
		db.Exec("UPDATE CONCEPTO SET prorrateo = true, basesac = true, tipoimpuestogananciasid = -2 WHERE ID = -12")
		db.Exec("UPDATE CONCEPTO SET prorrateo = false, tipoimpuestogananciasid = -10 WHERE ID IN (-18,-19)")
		db.Exec("UPDATE CONCEPTO SET prorrateo = false, tipoimpuestogananciasid = -11 WHERE ID = -20")

		if versionConceptoDB < 13 {
			db.Exec("update concepto set esremvariable = false where esremvariable is NULL")
			db.Exec("update concepto set esremvariable = true where id in (-5, -6)")
		}

		if versionConceptoDB < 14 {
			db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid, eseditable, tipocalculoautomaticoid, formulanombre, esremvariable) VALUES(-33, current_timestamp,'Medicina Prepaga', 'MEDICINA_PREPAGA',  '', 1, '',-46, false, -2,false, null, null, false, true, -2, true, -1, null, false)")
			db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid, eseditable, tipocalculoautomaticoid, formulanombre, esremvariable) VALUES(-34, current_timestamp,'Dias De Licencia', 'DIAS_DE_LICENCIA',  '', 1, '',-46, true, -3,true, null, null, false, true, -1, true, -1, null, false)")
		}

	}
	return err
}

func ObtenerConceptosPublicos(db *gorm.DB) error {
	//var conceptos []structConcepto.Concepto
	/*var tiposConcepto []structConcepto.Tipoconcepto
	var tiposCalculo []structConcepto.Tipodecalculo
	var tiposImpuestoGanancias []structConcepto.Tipoimpuestoganancias
	var tiposCalculoAutomatico []structConcepto.Tipocalculoautomatico*/

	//db_public := conexionBD.ObtenerDB("public")

	/*db_public.Find(&tiposConcepto)
	for _, tipoConcepto := range tiposConcepto {
		if err := db.Save(&tipoConcepto).Error; err != nil {
			return err
		}
	}

	db_public.Find(&tiposCalculo)
	for _, tipoCalculo := range tiposCalculo {
		if err := db.Save(&tipoCalculo).Error; err != nil {
			return err
		}
	}

	db_public.Find(&tiposImpuestoGanancias)
	for _, tipoImpuestoGanancias := range tiposImpuestoGanancias {
		if err := db.Save(&tipoImpuestoGanancias).Error; err != nil {
			return err
		}
	}

	db_public.Find(&tiposCalculoAutomatico)
	for _, tipoCalculoAutomatico := range tiposCalculoAutomatico {
		if err := db.Save(&tipoCalculoAutomatico).Error; err != nil {
			return err
		}
	}*/

	/*db_public.Find(&conceptos)
	for i := 0; i < len(conceptos); i++ {
		concepto := conceptos[i]
		if err := db.Save(&concepto).Error; err != nil {
			return err
		}
	}*/

	db.Exec("select ST_copy_concepto_public_privado()")

	return nil
}
