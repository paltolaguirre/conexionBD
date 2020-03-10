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
	db.Model(&structConcepto.Concepto{}).AddForeignKey("Formulanombre", "Function(Name)", "RESTRICT", "RESTRICT")

	versionConceptoDB := ObtenerVersionConceptoDB(db)

	if versionConceptoDB < 9 {
		err = db.Exec("update concepto set tipocalculoautomaticoid = -1 where tipodecalculoid is null").Error
		err = db.Exec("update concepto set tipocalculoautomaticoid = -2 where tipodecalculoid is not null").Error
	}
	return err
}

func AutomigrateConceptoTablasPublicas(db *gorm.DB) error {
	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structConcepto.Tipocalculoautomatico{}, &structConcepto.Tipoconcepto{}, &structConcepto.Tipodecalculo{}, &structConcepto.Tipoimpuestoganancias{}, &structConcepto.Concepto{}).Error
	if err == nil {

		versionConceptoDB := ObtenerVersionConceptoDB(db)

		db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid) VALUES(-29, current_timestamp,'Impuesto a las Ganancias', 'IMPUESTO_GANANCIAS',  '', 1, '',-46, true, -4,false, null, null, false, false, -1)")
		db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid) VALUES(-30, current_timestamp,'Impuesto a las Ganancias (Devolución)', 'IMPUESTO_GANANCIAS_DEVOLUCION',  '', 1, '',-46, true, -2,false, null, null, false, false, -1)")
		db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid) VALUES(-31, current_timestamp,'Cuota Sindical', 'CUOTA_SINDICAL',  '', 1, '',-46, true, -4,false, 2, -3, false, true, -12)")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -18")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -19")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -20")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -31")
		db.Exec("UPDATE CONCEPTO SET eseditable = false WHERE id in (-29, -30)")
		db.Exec("INSERT INTO TIPOCALCULOAUTOMATICO(id, created_at, nombre, codigo, descripcion, activo) VALUES(-1, current_timestamp, 'No Aplica', 'NO_APLICA', '', 1)")
		db.Exec("INSERT INTO TIPOCALCULOAUTOMATICO(id, created_at, nombre, codigo, descripcion, activo) VALUES(-2, current_timestamp, 'Porcentaje', 'PORCENTAJE', '', 1)")
		db.Exec("INSERT INTO TIPOCALCULOAUTOMATICO(id, created_at, nombre, codigo, descripcion, activo) VALUES(-3, current_timestamp, 'Formula', 'FORMULA', '', 1)")

		if versionConceptoDB < 9 {
			err = db.Exec("update concepto set tipocalculoautomaticoid = -1 where tipodecalculoid is null").Error
			err = db.Exec("update concepto set tipocalculoautomaticoid = -2 where tipodecalculoid is not null").Error
		}

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
