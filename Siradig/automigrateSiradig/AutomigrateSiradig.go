package automigrateSiradig

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Siradig/structSiradig"
)

func AutomigrateSiradigTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structSiradig.Detallecargofamiliarsiradig{}, &structSiradig.Importegananciasotroempleosiradig{}, &structSiradig.Deducciondesgravacionsiradig{}, &structSiradig.Retencionpercepcionsiradig{}, &structSiradig.Beneficiosiradig{}, &structSiradig.Ajustesiradig{}).Error
	if err == nil {
		db.Model(&structSiradig.Detallecargofamiliarsiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Detallecargofamiliarsiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

		db.Model(&structSiradig.Importegananciasotroempleosiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Importegananciasotroempleosiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

		db.Model(&structSiradig.Deducciondesgravacionsiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Deducciondesgravacionsiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

		db.Model(&structSiradig.Retencionpercepcionsiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Retencionpercepcionsiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

		db.Model(&structSiradig.Beneficiosiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Beneficiosiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

		db.Model(&structSiradig.Ajustesiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Ajustesiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

	}
	return err
}

func AutomigrateSiradigTablasPublicas(db *gorm.DB) error {
	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structSiradig.Siradigotipoimpuesto{}, &structSiradig.Siradigtipooperacion{}, &structSiradig.Siradigtipooperacion{}).Error
	if err == nil {
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES(current_timestamp,'Hijo Siradig','HIJO_SIRADIG','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Conyuge Siradig','CONYUGE_SIRADIG','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Cuota Médica Asistencial','CUOTA_MEDICA_ASISTENCIAL','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Primas de Seguro para el caso de muerte','PRIMAS_DE_SEGURO_PARA_EL_CASO_DE_MUERTE','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Primas de Seguro de Ahorro o Mixto','PRIMAS_DE_SEGURO_DE_AHORRO_O_MIXTO','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Aportes a Planes de Seguro de Retiro Privado','APORTES_A_PLANES_DE_SEGURO_DE_RETIRO_PRIVADO','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Donaciones','DONACIONES','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Intereses Préstamo Hipotecario','INTERESES_PRESTAMO_HIPOTECARIO','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Gastos de Sepelio','GASTOS_DE_SEPELIO','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Gastos Médicos y Paramédicos','GASTOS_MEDICOS_Y_PARAMEDICOS','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Gastos de Adquisición de indumentaria y equipamiento para uso exclusivo en el lugar de trabajo','GASTOS_ADQUISICION_INDUMENTARIA_Y_EQUIPAMIENTO_PARA_USO_EXCLUSIVO_EN_EL_LUGAR_DE_TRABAJO','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Alquiler de inmuebles destinados a casa habitación','ALQUILER_INMUEBLES_DESTINADOS_A_CASA_HABITACION','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Deducción del personal doméstico','DEDUCCION_DEL_PERSONAL_DOMESTICO','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Aportes a Sociedades de Garantía Recíproca ','APORTES_A_SOCIEDADES_DE_GARANTIA_RECIPROCA','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Vehículos de Viajantes y Corredores de Comercio','VEHICULOS_DE_VIAJANTES_Y_CORREDORES_DE_COMERCIOS','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Gastos de Representación e intereses de corredores y viajantes de comercio','GASTOS_DE_REPRESENTACION_E_INTERESES_DE_CORREDORES_Y_VIAJANTES_DE_COMERCIO','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Otras','OTRAS','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Adquisición de cuotapartes de FCI con fines de retiro','ADQUISICION_DE_CUOTAPARTES_DE_FCI_CON_FINES_DE_RETIRO','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Honorarios serv. Asistencia sanitaria, médica y paramédica','HONORARIOS_SERV_ASISTENCIA_SANITARIA_MEDICA_Y_PARAMEDICA','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Aportes Cajas complementarias / Fondos compensadores de prev. / Similares','APORTES_CAJAS_COMPLEMENTARIAS_FONDOS_COMPENSADORES_DE_PREV_SIMILARES','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Impuestos sobre créditos y débitos','IMPUESTOS_SOBRE_CREDITOS_Y_DEBITOS','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Percepciones / Retenciones aduaneras','PERCEPCIONES_RETENCIONES_ADUANERAS','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Pago a Cuenta - Resolucion General (AFIP) 3819/2015 - Cancelaciones en Efectivo','PAGO_A_CUENTA_RESOLUCION_GENERAL_AFIP_3819_2015_CANCELACIONES_EN_EFECTIVO','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Sr. TRABAJADOR: ¿Usted trabaja en Región Patagónica? (art. 1ro Ley 23.272)','SR_TRABAJADOR_USTED_TRABAJA_EN_REGION_PATAGONICA_ART_1RO_LEY_23272','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Sr. JUBILADO/PENSIONADO Y/O RETIRADO: ¿Usted vive en Región Patagónica? (art. 1ro Ley 23.272)','SR_JUBULADO_PENSIONADO_RETIRADO_USTED_VIVE_EN_REGION_PATAGONICA_ART_1RO_LEY_23272','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Sr. JUBILADO/PENSIONADO Y/O RETIRADO: ¿Percibe otros ingresos por monotributo/ relación de dependencia/ participación en sociedades/ actividad autónoma/ intereses de plazos fijos/ etc.','SR_JUBILADO_PENSIONADO_RETIRADO_PERCIBE_OTROS_INGRESOS_POR_MONOTRIBUTO_RELACION_DE_DEPENDENCIA_PARTICIPACION_EN_SOCIEDADES_ACTIVIDAD_AUTONOMA_INTERESES_DE_PLAZOS_FIJOS_ETC','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Sr. JUBILADO/PENSIONADO Y/O RETIRADO: ¿Tributó Bienes Personales en el último período fiscal anterior al que está declarando?','','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Sr. JUBILADO/PENSIONADO Y/O RETIRADO: ¿Tiene más bienes, por los que tributó Bienes Personales en el último período fiscal anterior al que está declarando, aparte de su casa-habitación?','SR_JUBILADO_PENSIONADO_RETIRADO_TIENE_MAS_BIENES_POR_LOS_QUE_TRIBUTO_BIENES_PERSONALES_EN_EL_ULTIMO_PERIODO_FISCAL_ANTERIOR_AL_QUE_ESTA_DECLARANDO_APARTE_DE_SU_CASA_HABITACION','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Ajustes de Haberes Retroactivos de Otros Empleadores Percibidos en el Período Declarado','AJUSTES_DE_HABERES_RETROACTIVOS_DE_OTROS_EMPLEADORES_PERCIBIDOS_EN_EL_PERIODO_DECLARADO','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOGRILLA(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Reintegro de Aportes de Socios Protectores a Sociedades de Garantía Recíproca','REINTEGRO_DE_APORTES_DE_SOCIOS_PROTECTORES_A_SOCIEDADES_DE_GARANTIA_RECIPROCA','',1)")

		db.Exec("UPDATE public.SIRADIGTIPOGRILLA SET id = -id;")

		db.Exec("INSERT INTO SIRADIGTIPOIMPUESTO(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Impuestos sobre créditos y débitos en cuenta Bancaria','IMPUESTOS_SOBRE_CREDITOS_Y_DEBITOS_EN_CUENTA_BANCARIA','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOIMPUESTO(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Impuestos sobre movimientos de fondos propios o terceros','IMPUESTOS_SOBRE_MOVIMIENTOS_DE_FONDOS_PROPIOS_O_TERCEROS','',1)")

		db.Exec("UPDATE public.SIRADIGTIPOIMPUESTO SET id = -id;")

		db.Exec("INSERT INTO SIRADIGTIPOOPERACION(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Operaciones de adquisición de servicios en el exterior contratados a través de agencias de viajes y turismo - mayoristas y/o minoristas - del pais, que se cancelen mediante el pago en efectivo (R.G. 3819 - Art. 1° inciso a) ','OPERACIONES_DE_ADQUISICION_DE_SERVICIOS_EN_EL_EXTERIOR_CONTRATADOS_A_TRAVES_DE_AGENCIAS_DE_VIAJES_Y_TURISMO_QUE_SE_CANCELEN_MEDIANTE_EL_PAGO_EN_EFECTIVO_RG_3819_ART_1_INCISO_A','',1)")
		db.Exec("INSERT INTO SIRADIGTIPOOPERACION(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Operaciones de adquisición de servicios de transporte terrestre, aéreo y por vía acuática, de pasajeros con destino fuera del pais, que se cancelen mediante el pago en efectivo (R,G, 3819 - Art. 1° inciso b)','OPERACIONES_DE_ADQUISICION_DE_SERVICIOS_DE_TRANSPORTE_TERRESTRE_AEREO_Y_POR_VIA_ACUATICA_DE_PASAJEROS_CON_DESTINO_FUERA_DEL_PAIS_QUE_SE_CANCELEN_MEDIANTE_EL_PAGO_EN_EFECTIVO_RG_3819_ART_1_INCISO_B','',1)")

		db.Exec("UPDATE public.SIRADIGTIPOOPERACION SET id = -id;")

	}

	return err
}
