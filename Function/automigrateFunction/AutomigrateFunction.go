package automigrateFunction

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Function/structFunction"
)

func AutomigrateFunctionTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structFunction.Invoke{}, &structFunction.Value{}, &structFunction.Param{}, &structFunction.Function{}).Error
	if err == nil {
		db.Model(&structFunction.Param{}).AddForeignKey("functionname", "function(name)", "CASCADE", "CASCADE")
		db.Model(&structFunction.Function{}).AddForeignKey("valueid", "value(id)", "CASCADE", "CASCADE")
		db.Model(&structFunction.Invoke{}).AddForeignKey("functionname", "function(name)", "CASCADE", "CASCADE")
		db.Model(&structFunction.Value{}).AddForeignKey("valueinvokeid", "invoke(id)", "CASCADE", "CASCADE")
	}
	return err

}

func AutomigrateFunctionTablasPublicas(db *gorm.DB) error {
	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structFunction.Value{}, &structFunction.Invoke{}, &structFunction.Param{}, &structFunction.Function{}).Error
	if err == nil {
		versionFunctionDB := ObtenerVersionFunctionDB(db)
		if versionFunctionDB < 1 {
			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-1,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('GetParamValue', current_timestamp, 'Obtiene el valor de un parametro de la formula', 'primitive', 'internal', 'public', 'number', -1)")
			db.Exec("INSERT INTO param(id, created_at, name, type, functionname) VALUES(-1, current_timestamp, 'paramName', 'string', 'GetParamValue')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-2,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Sum', current_timestamp, 'Dado dos valores obtiene la suma de ambos', 'primitive', 'operator', 'public', 'number', -2)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-2,current_timestamp,'val1','number','Sum')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-3,current_timestamp,'val2','number','Sum')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-3,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('TotalImporteRemunerativo', current_timestamp, 'Dada una liquidacion obtiene la suma total de importes remunerativos de la misma', 'primitive', 'helper', 'public', 'number', -3)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-4,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('ImpuestoALasGanancias', current_timestamp, '', 'primitive', 'helper', 'public', 'number', -4)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-5,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('ImpuestoALasGananciasDevolucion', current_timestamp, '', 'primitive', 'helper', 'public', 'number', -5)")
		}
		
		if versionFunctionDB < 2 {

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-6,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Sueldo', current_timestamp, 'Sueldo ingresado en el campo Remuneración del Legajo', 'primitive', 'helper', 'public', 'number', -6)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-7,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('HorasMensuales', current_timestamp, 'Horas ingresadas en el campo Horas Mensuales Normales del legajo', 'primitive', 'helper', 'public', 'number', -7)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-8,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasMesTrabajadosFechaLiquidacion', current_timestamp, 'Cantidad de dias desde el primero del mes hasta la fecha de la liquidación', 'primitive', 'helper', 'public', 'number', -8)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-9,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasMesTrabajadosFechaPeriodo', current_timestamp, 'Cantidad de dias desde el primero del mes hasta el último día del período liquidado', 'primitive', 'helper', 'public', 'number', -9)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-10,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemRemunerativaSemestre', current_timestamp, 'Mejor Remuneración del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre), comparando grillas de (Remunerativo - Desacuentos)', 'primitive', 'helper', 'public', 'number', -10)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-11,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasSemTrabajados', current_timestamp, 'Dias Trabajados en el Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre)', 'primitive', 'helper', 'public', 'number', -11)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-12,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemNoRemunerativaSemestre', current_timestamp, 'Mejor Remuneración No Remunerativa del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre)', 'primitive', 'helper', 'public', 'number', -12)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-13,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('ValorDiasVacaciones', current_timestamp, 'Valor de los días correspondientes a las Vacaciones', 'primitive', 'helper', 'public', 'number', -13)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-14, current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('TotalHaberesNoRemunerativosMensual', current_timestamp, 'Total de conceptos no remunerativos de la liquidación', 'primitive', 'helper', 'public', 'number', -14)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-15,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('TotalAportesPatronalesMensual', current_timestamp, 'Total de conceptos de aportes patronales de la liquidación', 'primitive', 'helper', 'public', 'number', -15)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-16,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('TotalRetencionesMensual', current_timestamp, 'Total de conceptos de retenciones de la liquidación', 'primitive', 'helper', 'public', 'number', -16)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-17,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('CantidadMesesTrabajados', current_timestamp, 'La cantidad de meses trabajados desde  Legajo > Fecha de Ingreso  hasta la Fecha de la liquidación', 'primitive', 'helper', 'public', 'number', -17)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-18,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasLicenciaMes', current_timestamp, 'Cantidad de días de licencias del mes de liquidación', 'primitive', 'helper', 'public', 'number', -18)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-19,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasLicenciaSemestre', current_timestamp, 'Cantidad de días de licencias del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre)', 'primitive', 'helper', 'public', 'number', -19)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-20,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('TotalDescuentosMensual', current_timestamp, 'Total de conceptos de descuentos de la liquidación', 'primitive', 'helper', 'public', 'number', -20)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-21,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemNormalYHabitualSemestre', current_timestamp, 'Mejor remuneración normal y habitual del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre)', 'primitive', 'helper', 'public', 'number', -21)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-22,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasEfectivamenteTrabajadosSemestre', current_timestamp, 'Cantidad de días efectivamente trabajados en el Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre). Días hábiles - Días Faltas injustificadas - Días enfermedad - Días accidente - Días de Licencia', 'primitive', 'helper', 'public', 'number', -22)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-23,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('FechadeIngreso', current_timestamp, 'Fecha de ingreso del Legajo', 'primitive', 'helper', 'public', 'Time', -23)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-24,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('FechadeLiquidacion', current_timestamp, 'Fecha de la liquidación donde se esta utilizando el concepto', 'primitive', 'helper', 'public', 'Time', -24)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-25,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('FecIngHASTAFecLiq', current_timestamp, 'Fecha de Ingreso hasta Fecha de Liquidacion expresada en años', 'primitive', 'helper', 'public', 'number', -25)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-26,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('AntiguedadResto', current_timestamp, ' ( (Dias desde la Fecha Ingreso hasta la Fecha Liquidación) / 365 ) - antiguedad', 'primitive', 'helper', 'public', 'number', -26)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-27,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemTotalSemestre', current_timestamp, 'Mejor remuneración total del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre). Comparando las grillas de Remunerativo - Descuentos + No Remunerativo', 'primitive', 'helper', 'public', 'number', -27)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-28,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemTotalAnual', current_timestamp, 'Mejor remuneración total del año. Comparando las grillas de Remunerativo - Descuentos + No Remunerativo', 'primitive', 'helper', 'public', 'number', -28)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-29,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemTotalSinRemVarSemestre', current_timestamp, 'Mejor remuneración total, sin incluir Remuneraciones Variables del Semestre. Comparando las grillas de Remunerativo - Descuentos + No Remunerativo', 'primitive', 'helper', 'public', 'number', -29)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-30,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemTotalSinRemVarAnual', current_timestamp, 'Mejor remuneración total, sin remuneraciones variables del año liquidado. Comparando las grillas de Remunerativo - Descuentos + No Remunerativo', 'primitive', 'helper', 'public', 'number', -30)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-31,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('PromRemVariablesSemestre', current_timestamp, 'Promedio remuneraciones variables del Semestre', 'primitive', 'helper', 'public', 'number', -31)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-32,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('PromRemVariablesAnual', current_timestamp, 'Promedio remuneraciones variables del Año Liquidado', 'primitive', 'helper', 'public', 'number', -32)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-33,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemRemunerativaBaseSACSemestre', current_timestamp, 'Mejor Remuneración del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre), comparando grillas de Remunerativo - Descuentos solo de los conceptos que tienen configurado BASE_SAC = SI', 'primitive', 'helper', 'public', 'number', -33)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-34,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('GetConceptValue', current_timestamp, 'Obtiene el valor de un concepto', 'primitive', 'internal', 'public', 'number', -34)")
			db.Exec("INSERT INTO param(id, created_at, name, type, functionname) VALUES(-4, current_timestamp, 'conceptid', 'number', 'GetConceptValue')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-35,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('If', current_timestamp, 'Retorna el primer valor si la condicion es verdadera y el segundo valor si la condicion es falsa', 'primitive', 'internal', 'public', 'number', -35)")
			db.Exec("INSERT INTO param(id, created_at, name, type, functionname) VALUES(-5, current_timestamp, 'condicion', 'bool', 'If')")
			db.Exec("INSERT INTO param(id, created_at, name, type, functionname) VALUES(-6, current_timestamp, 'valueTrue', 'number', 'If')")
			db.Exec("INSERT INTO param(id, created_at, name, type, functionname) VALUES(-7, current_timestamp, 'valueFalse', 'number', 'If')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-36,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Equality', current_timestamp, 'Dado dos valores retorna si son iguales', 'primitive', 'operator', 'public', 'bool', -36)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-8,current_timestamp,'val1','number','Equality')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-9,current_timestamp,'val2','number','Equality')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-37,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Inequality', current_timestamp, 'Dado dos valores retorna si son distintos', 'primitive', 'operator', 'public', 'bool', -37)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-10,current_timestamp,'val1','number','Inequality')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-11,current_timestamp,'val2','number','Inequality')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-38,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Greater', current_timestamp, 'Dado dos valores retorna si el primero es mayor al segundo', 'primitive', 'operator', 'public', 'bool', -38)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-12,current_timestamp,'val1','number','Greater')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-13,current_timestamp,'val2','number','Greater')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-39,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('GreaterEqual', current_timestamp, 'Dado dos valores retorna si el primero es mayor o igual al segundo', 'primitive', 'operator', 'public', 'bool', -39)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-14,current_timestamp,'val1','number','GreaterEqual')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-15,current_timestamp,'val2','number','GreaterEqual')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-40,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Less', current_timestamp, 'Dado dos valores retorna si el primero es menor al segundo', 'primitive', 'operator', 'public', 'bool', -40)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-16,current_timestamp,'val1','number','Less')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-17,current_timestamp,'val2','number','Less')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-41,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('LessEqual', current_timestamp, 'Dado dos valores retorna si el primero es menor o igual al segundo', 'primitive', 'operator', 'public', 'bool', -41)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-14,current_timestamp,'val1','number','LessEqual')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-15,current_timestamp,'val2','number','LessEqual')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-42,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Not', current_timestamp, 'Dada un valor de verdad, devuelve el contrario', 'primitive', 'operator', 'public', 'bool', -42)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-16,current_timestamp,'val1','bool','Not')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-43,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('And', current_timestamp, 'Dados dos valores de verdad, devuelve el valor de verdad de la conjuncion (Solo es verdadero si ambos son verdaderos)', 'primitive', 'operator', 'public', 'bool', -43)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-17,current_timestamp,'val1','bool','And')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-18,current_timestamp,'val2','bool','And')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-44,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Or', current_timestamp, 'Dados dos valores de verdad, devuelve el valor de verdad de la disyuncion no excluyente (Solo es verdadero si alguno es verdadero)', 'primitive', 'operator', 'public', 'bool', -44)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-19,current_timestamp,'val1','bool','Or')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-20,current_timestamp,'val2','bool','Or')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-45,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Percent', current_timestamp, 'Dado un numero y un porcentaje, devuelve el numero que representa ese porcentaje', 'primitive', 'operator', 'public', 'number', -45)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-21,current_timestamp,'val','number','Percent')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-22,current_timestamp,'percent','number','Percent')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-46,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Diff', current_timestamp, 'Devuelve el primer numero menos el segundo numero', 'primitive', 'operator', 'public', 'number', -46)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-23,current_timestamp,'val1','number','Diff')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-24,current_timestamp,'val2','number','Diff')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-47,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Div', current_timestamp, 'Devuelve el primer numero dividido el segundo numero', 'primitive', 'operator', 'public', 'number', -47)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-25,current_timestamp,'val1','number','Div')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-26,current_timestamp,'val2','number','Div')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-48,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Multi', current_timestamp, 'Devuelve el primer numero multiplicado por el segundo numero', 'primitive', 'operator', 'public', 'number', -48)")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-27,current_timestamp,'val','number','Multi')")
			db.Exec("INSERT INTO param(id,created_at, name, type, functionname) VALUES(-28,current_timestamp,'percent','number','Multi')")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-49,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Jornal', current_timestamp, 'Variable Sueldo / 30', 'primitive', 'helper', 'public', 'number', -49)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-50,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('ValorHora', current_timestamp, 'Variables Sueldo / HorasMensuales', 'primitive', 'helper', 'public', 'number', -50)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-51,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('HoraExtra50', current_timestamp, 'Variable ValorHora * 1,5 * Cantidad (será el campo Cantidad Ingresado en una Novedad)', 'primitive', 'helper', 'public', 'number', -51)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-52,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('HoraExtra100', current_timestamp, 'Variable ValorHora * 2 * Cantidad (será el campo Cantidad Ingresado en una Novedad)', 'primitive', 'helper', 'public', 'number', -52)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-53,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Sac', current_timestamp, 'Variable (MejorRemRemunerativaBaseSACSemestre / 2) * DiasSemTrabajados / 180', 'primitive', 'helper', 'public', 'number', -53)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-54,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('SacNoRemunerativo', current_timestamp, 'Variable (MejorRemNoRemunerativaSemestre / 2) * DiasSemTrabajados / 180', 'primitive', 'helper', 'public', 'number', -54)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-55,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Antiguedad', current_timestamp, 'Variable (MejorRemNoRemunerativaSemestre / 2) * DiasSemTrabajados / 180', 'primitive', 'helper', 'public', 'number', -55)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-56,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasFaltasInjustificadas', current_timestamp, '', 'primitive', 'helper', 'public', 'number', -56)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-57,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Vacaciones', current_timestamp, 'Si la antiguedad es menor a 5 entonces le corresponden 14 dias de vacaciones. Entre 5 y 10 le corresponden 21 dias de vacaciones. Entre 10 y 15 le corresponden 28 dias de vacaciones. Mayor a 15 le corresponden 35 dias de vacaciones.', 'primitive', 'helper', 'public', 'number', -57)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-58,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Preaviso', current_timestamp, '', 'primitive', 'helper', 'public', 'number', -58)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-59,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('SacSinPreaviso', current_timestamp, 'SAC sobre Preaviso', 'primitive', 'helper', 'public', 'number', -59)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-60,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('IntegracionMesDespido', current_timestamp, '( 30 - DiasMesTrabajados ) * Jornal', 'primitive', 'helper', 'public', 'number', -60)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-61,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('SacSinImd', current_timestamp, 'Integración mes Despido / 12', 'primitive', 'helper', 'public', 'number', -61)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-62,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('IndemnizacionPorDespido', current_timestamp, '', 'primitive', 'helper', 'public', 'number', -62)")

			db.Exec("INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-63,current_timestamp,'return',0,'',false,null,0)")
			db.Exec("INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('CantidadSueldos', current_timestamp, '', 'primitive', 'helper', 'public', 'number', -63)")



		}
	}
	return err
}

func ObtenerFormulasPublicas(db *gorm.DB) error {
	var formulas []structFunction.Function

	db_public := conexionBD.ObtenerDB("public")

	db_public.Set("gorm:auto_preload", true).Find(&formulas)
	for i := 0; i < len(formulas); i++ {
		formula := formulas[i]

		params := formula.Params

		formula.Params = nil
		if err := db.Save(&formula.Value).Error; err != nil {
			return err
		}
		if err := db.Save(&formula).Error; err != nil {
			return err
		}
		for _, param := range params {
			if err := db.Save(&param).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
