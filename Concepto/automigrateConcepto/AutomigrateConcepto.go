package automigrateConcepto

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Concepto/structConcepto"
	"github.com/xubiosueldos/conexionBD/versiondbmicroservicio"
)

type MicroservicioConcepto struct {
}

func (*MicroservicioConcepto) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(ObtenerVersionConceptoConfiguracion(), ObtenerVersionConceptoDB(db))
}

func (*MicroservicioConcepto) AutomigrarPublic(db *gorm.DB) error {
	err := AutomigrateConceptoTablasPublicas(db)
	return err
}

func (*MicroservicioConcepto) AutomigrarPrivate(db *gorm.DB) error {
	if err := AutomigrateConceptoTablasPrivadas(db); err != nil {
		return err
	} else {
		if err = ObtenerConceptosPublicos(db); err != nil {
			return err
		}
	}
	return nil
}

func (*MicroservicioConcepto) ActualizarVersion(db *gorm.DB) {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(ObtenerVersionConceptoConfiguracion(), Concepto, db)
}

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
	err := db.AutoMigrate(&structConcepto.Tipocalculoautomatico{}, &structConcepto.Tipoconcepto{}, &structConcepto.Tipodecalculo{}, &structConcepto.Tipoimpuestoganancias{}, &structConcepto.Conceptoafip{}, &structConcepto.Concepto{}).Error
	if err == nil {

		versionConceptoDB := ObtenerVersionConceptoDB(db)

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

		db.Exec("INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) VALUES(-16,current_timestamp,'SAC','SAC','',1, false, false, false, true, false)")
		db.Exec("UPDATE CONCEPTO SET tipoimpuestogananciasid = -16 WHERE id = -2")

		if versionConceptoDB < 10 {
			err = db.Exec("update concepto set tipocalculoautomaticoid = -1 where tipodecalculoid is null").Error
			err = db.Exec("update concepto set tipocalculoautomaticoid = -2 where tipodecalculoid is not null").Error

			err = db.Exec("update concepto set formulanombre = 'ImpuestoALasGanancias', tipocalculoautomaticoid = -3 where id = -29").Error
			err = db.Exec("update concepto set formulanombre = 'ImpuestoALasGananciasDevolucion', tipocalculoautomaticoid = -3 where id = -30").Error

			err = db.Exec("update concepto set eseditable = false where tipocalculoautomaticoid != -1").Error
			err = db.Exec("update concepto set eseditable = true where tipocalculoautomaticoid = -1").Error
		}

		db.Exec("UPDATE CONCEPTO SET prorrateo = false, basesac = true, tipoimpuestogananciasid = -1 WHERE ID IN (-1,-3,-4,-15,-16,-17)")
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

		if versionConceptoDB < 16 {
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -1 WHERE id IN (-1,-3,-4,-15,-16,-17)")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -11 WHERE id = -2")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -16 WHERE id = -5")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -17 WHERE id = -6")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -56 WHERE id = -7")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -57 WHERE id = -8")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -59 WHERE id IN (-9,-14)")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -60 WHERE id = -10")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -54 WHERE id = -11")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -52 WHERE id = -12")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -58 WHERE id = -13")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -65 WHERE id = -18")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -66 WHERE id = -19")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -67 WHERE id = -20")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -73 WHERE id = -29")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -69 WHERE id = -31")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -63 WHERE id = -30")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -21 WHERE id = -32")
			db.Exec("UPDATE CONCEPTO SET conceptoafipid = -6 WHERE id = -34")
			db.Exec("UPDATE CONCEPTO SET marcarepeticion = true, aportesipa = true, contribucionsipa = true, aportesinssjyp = true, contribucionesinssjyp = true, aportesobrasocial = true, contribucionesobrasocial = true, aportesfondosolidario = true, contribucionesfondosolidario = true, aportesrenatea = true, contribucionesrenatea = true, asignacionesfamiliares = true, contribucionesfondonacional = true, contribucionesleyriesgo = true, aportesregimenesdiferenciales = false, aportesregimenesespeciales = false WHERE ID IN (-1,-2,-3,-4,-5,-6,-15,-16,-17,-32,-34)")
			db.Exec("UPDATE CONCEPTO SET marcarepeticion = true, aportesipa = false, contribucionsipa = false, aportesinssjyp = false, contribucionesinssjyp = false, aportesobrasocial = false, contribucionesobrasocial = false, aportesfondosolidario = false, contribucionesfondosolidario = false, aportesrenatea = false, contribucionesrenatea = false, asignacionesfamiliares = false, contribucionesfondonacional = false, contribucionesleyriesgo = true, aportesregimenesdiferenciales = false, aportesregimenesespeciales = false WHERE ID IN (-7,-8,-9,-10,-11,-12,-13,-14)")
			db.Exec("UPDATE CONCEPTO SET marcarepeticion = true, aportesipa = false, contribucionsipa = false, aportesinssjyp = false, contribucionesinssjyp = false, aportesobrasocial = false, contribucionesobrasocial = false, aportesfondosolidario = false, contribucionesfondosolidario = false, aportesrenatea = false, contribucionesrenatea = false, asignacionesfamiliares = false, contribucionesfondonacional = false, contribucionesleyriesgo = false, aportesregimenesdiferenciales = false, aportesregimenesespeciales = false WHERE ID IN (-18,-19,-20,-29,-30,-31)")
		}

		if versionConceptoDB < 17 {
			db.Exec("update concepto set formulanombre = 'Sac', tipocalculoautomaticoid = -3 where nombre = 'Sueldo Anual Complementario'")
			db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid, eseditable, tipocalculoautomaticoid, formulanombre, esremvariable, conceptoafipid, marcarepeticion, aportesipa, contribucionsipa, aportesinssjyp, contribucionesinssjyp, aportesobrasocial, contribucionesobrasocial, aportesfondosolidario, contribucionesfondosolidario, aportesrenatea, contribucionesrenatea, asignacionesfamiliares, contribucionesfondonacional, contribucionesleyriesgo, aportesregimenesdiferenciales, aportesregimenesespeciales, codigointerno) VALUES(-35, current_timestamp,'Sueldo Anual Complementario No Remunerativo', 'SUELDO_ANUAL_COMPLEMENTARIO_NO_REMUNERATIVO',  '', 1, '',-46, true, -2,false, null, null, null, null, -16, true, -3, 'SacNoRemunerativo', false, -11, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, 18)")
			db.Exec("update concepto set formulanombre = 'Vacaciones', tipocalculoautomaticoid = -3 where nombre = 'Vacaciones'")
			db.Exec("update concepto set formulanombre = 'Preaviso', tipocalculoautomaticoid = -3 where nombre = 'Preaviso'")
			db.Exec("update concepto set formulanombre = 'SacSinPreaviso', tipocalculoautomaticoid = -3 where nombre = 'SAC sobre Preaviso'")
			db.Exec("update concepto set formulanombre = 'IntegracionMesDespido', tipocalculoautomaticoid = -3 where nombre = 'Integración Mes de despido'")
		}
	}
	return err
}

func ObtenerConceptosPublicos(db *gorm.DB) error {

	versionConceptoDB := ObtenerVersionConceptoDB(db)

	db.Exec("select ST_copy_concepto_public_privado()")

	if versionConceptoDB < 16 {
		db.Exec("ALTER SEQUENCE concepto_codigointerno_seq RESTART with 1000")
		db.Exec("UPDATE CONCEPTO SET codigointerno = nextval('concepto_codigointerno_seq') WHERE id > 0")
	}

	return nil
}
