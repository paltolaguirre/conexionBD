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
	err := db.AutoMigrate(&structConcepto.Concepto{}, &structConcepto.Tipoconcepto{}, &structConcepto.Tipodecalculo{}).Error
	if err == nil {
		var tipodeCalculo structConcepto.Tipodecalculo
		db.Raw("SELECT * FROM TIPODECALCULO").Scan(&tipodeCalculo)
		if tipodeCalculo.ID == 0 {
			db.Exec("INSERT INTO TIPODECALCULO(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Remunerativos','REMUNERATIVOS','',1),(current_timestamp,'No Remunerativos','NO_REMUNERATIVOS','',1),(current_timestamp,'Remunerativos - Descuentos','REMUNERATIVOS_DESCUENTOS','',1),(current_timestamp,'Remunerativos + No Remunerativos','REMUNERATIVOS_NO_REMUNERATIVOS','',1),(current_timestamp,'Remunerativos + No Remunerativos - Descuentos','REMUNERATIVOS_NO_REMUNERATIVOS_DESCUENTOS','',1);")
			db.Exec("UPDATE public.tipodecalculo SET id = -id;")
		}

		db.Raw("UPDATE CONCEPTO SET porcentaje = 11.0000, tipodecalculoid = -3 WHERE ID = -18")
		db.Raw("UPDATE CONCEPTO SET porcentaje = 3.0000, tipodecalculoid = -3 WHERE ID = -19")
		db.Raw("UPDATE CONCEPTO SET porcentaje = 3.0000, tipodecalculoid = -3 WHERE ID = -20")

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
