package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Importegananciasotroempleosiradig struct {
	structGormModel.GormModel
	Siradigid                                 *int       `json:"siradigid"`
	Mes                                       *time.Time `json:"mes"`
	Importegananciasbrutas                    *float64   `json:"importegananciasbrutas"  sql:"type:decimal(19,4);"`
	Aporteseguridadsocial                     *float64   `json:"aporteseguridadsocial"  sql:"type:decimal(19,4);"`
	Aporteobrasocial                          *float64   `json:"aporteobrasocial"  sql:"type:decimal(19,4);"`
	Aportesindical                            *float64   `json:"aportesindical"  sql:"type:decimal(19,4);"`
	Importeretribucionesnohabituales          *float64   `json:"importeretribucionesnohabituales"  sql:"type:decimal(19,4);"`
	Importeretencionesgananciassufridas       *float64   `json:"importeretencionesgananciassufridas"  sql:"type:decimal(19,4);"`
	Ajustes                                   *float64   `json:"ajustes"  sql:"type:decimal(19,4);"`
	Importeconceptosexentos                   *float64   `json:"importeconceptosexentos"  sql:"type:decimal(19,4);"`
	Sac                                       *float64   `json:"sac"  sql:"type:decimal(19,4);"`
	Importehorasextrasgravadas                *float64   `json:"importehorasextrasgravadas"  sql:"type:decimal(19,4);"`
	Importehorasextrasexentas                 *float64   `json:"importehorasextrasexentas"  sql:"type:decimal(19,4);"`
	Materialdidactico                         *float64   `json:"materialdidactico"  sql:"type:decimal(19,4);"`
	Gastosmovilidad                           *float64   `json:"gastosmovilidad"  sql:"type:decimal(19,4);"`
	Ajusteperiodoanteriorremgravada           *float64   `json:"ajusteperiodoanteriorremgravada"  sql:"type:decimal(19,4);"`
	Ajusteperiodoanteriorremexentanoalcanzada *float64   `json:"ajusteperiodoanteriorremexentanoalcanzada"  sql:"type:decimal(19,4);"`
}
