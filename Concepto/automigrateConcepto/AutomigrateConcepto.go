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

		db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid) VALUES(-29, current_timestamp,'Impuesto a las Ganancias', 'IMPUESTO_GANANCIAS',  '', 1, '',-46, true, -4,false, null, null, false, false, -1)")
		db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid) VALUES(-30, current_timestamp,'Impuesto a las Ganancias (Devolución)', 'IMPUESTO_GANANCIAS_DEVOLUCION',  '', 1, '',-46, true, -2,false, null, null, false, false, -1)")
		db.Exec("INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid) VALUES(-31, current_timestamp,'Cuota Sindical', 'CUOTA_SINDICAL',  '', 1, '',-46, true, -4,false, 2, -3, false, true, -12)")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -18")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -19")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -20")
		db.Exec("UPDATE CONCEPTO SET basesac = true WHERE id = -31")
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
