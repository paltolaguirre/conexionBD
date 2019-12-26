package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Ajustesiradig struct {
	structGormModel.GormModel
	Siradigtipogrilla            Siradigtipogrilla `json:"siradigtipogrilla" gorm:"ForeignKey:Siradigtipogrillaid;association_foreignkey:ID;association_autoupdate:false"`
	Siradigtipogrillaid          *int              `json:"siradigtipogrillaid" sql:"type:int REFERENCES Siradigtipogrilla(ID)"`
	Siradigid                    int               `json:"siradigid"`
	Cuit                         string            `json:"cuir"`
	Razonsocial                  string            `json:"razonsocial"`
	Año                          *time.Time        `json:"año"`
	Montoretroactivocobrado      *float64          `json:"montoretroactivocobrado"  sql:"type:decimal(19,4);"`
	Cumplesegundoparrafoley24467 bool              `json:"cumplesegundoparrafoley24467"`
	Montoreintegrar              *float64          `json:"montoreintegrar"  sql:"type:decimal(19,4);"`
	Cumpletercerparrafoley24467  bool              `json:"cumpletercerparrafoley24467"`
	Montoreintegrar3             *float64          `json:"montoreintegrar3"  sql:"type:decimal(19,4);"`
}
