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
	err := db.AutoMigrate(&structConcepto.Concepto{}, &structConcepto.Tipoconcepto{}).Error
	if err == nil {
		var tipoConcepto structConcepto.Tipoconcepto
		db.Raw("SELECT * FROM TIPOCONCEPTO").Scan(&tipoConcepto)
		if tipoConcepto.ID == 0 {
			db.Exec("INSERT INTO TIPOCONCEPTO(created_at, nombre, codigo, descripcion, activo) VALUES (current_timestamp,'Importe Remunerativo','IMPORTE_REMUNERATIVO','',1),(current_timestamp,'Importe No Remunerativo','IMPORTE_NO_REMUNERATIVO','',1),(current_timestamp,'Descuento','DESCUENTO','',1),(current_timestamp,'Retención','RETENCION','',1),(current_timestamp,'Aporte Patronal','APORTE_PATRONAL','',1);")
			db.Exec("UPDATE public.tipoconcepto SET id = -id;")
		}
		db.Exec("UPDATE concepto SET tipoconceptoid = -1 ,esnovedad = false WHERE id IN (-1,-2,-3,-4)")
		db.Exec("UPDATE concepto SET tipoconceptoid = -1 ,esnovedad = true WHERE id IN (-5,-6)")
		db.Exec("UPDATE concepto SET tipoconceptoid = -2 ,esnovedad = false WHERE id IN (-7,-8,-9,-10,-11,-12,-13,-14)")
		db.Exec("UPDATE concepto SET tipoconceptoid = -3 ,esnovedad = true WHERE id IN (-15,-16,-17)")
		db.Exec("UPDATE concepto SET tipoconceptoid = -4 ,esnovedad = false WHERE id IN (-18,-19,-20)")
		db.Exec("UPDATE concepto SET tipoconceptoid = -5 ,esnovedad = false WHERE id IN (-21,-22,-23,-24,-25,-26,-27,-28)")

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
