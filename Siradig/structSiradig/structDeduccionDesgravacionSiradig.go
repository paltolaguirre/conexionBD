package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Deducciondesgravacionsiradig struct {
	structGormModel.GormModel
	Siradigtipogrilla      Siradigtipogrilla `json:"siradigtipogrilla" gorm:"ForeignKey:Siradigtipogrillaid;association_foreignkey:ID;association_autoupdate:false"`
	Siradigtipogrillaid    *int              `json:"siradigtipogrillaid" sql:"type:int REFERENCES Siradigtipogrilla(ID)"`
	Siradigid              int               `json:"siradigid"`
	Mes                    *time.Time        `json:"mes"`
	Meshasta               *time.Time        `json:"meshasta"`
	Importe                *float64          `json:"importe"  sql:"type:decimal(19,4);"`
	Descripcion            string            `json:"descripcion"`
	Comprobante            string            `json:"comprobante"`
	Contribucion           *float64          `json:"contribucion"  sql:"type:decimal(19,4);"`
	Retribucion            *float64          `json:"retribucion"  sql:"type:decimal(19,4);"`
	Cuit                   string            `json:"cuit"`
	Empresa                string            `json:"empresa"`
	Montocapitalaporte     *float64          `json:"montocapitalaporte"  sql:"type:decimal(19,4);"`
	Montofondoriesgoaporte *float64          `json:"montofondoriesgoaporte"  sql:"type:decimal(19,4);"`
	Valor                  *float64          `json:"valor"  sql:"type:decimal(19,4);"`
	Porcentajeafectacion   string            `json:"porcentajeafectacion"`
	Amortizacionperiodo    string            `json:"amortizacionperiodo"`
}
