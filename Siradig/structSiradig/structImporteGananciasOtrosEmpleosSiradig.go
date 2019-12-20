package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Importegananciasotrosempleossiradig struct {
	structGormModel.GormModel
	Mes                                 *time.Time `json:"mes"`
	Importegananciasbrutas              *float64   `json:"importegananciasbrutas"`
	Aporteseguridadsocial               *float64   `json:"aporteseguridadsocial"`
	Aporteobrasocial                    *float64   `json:"aporteobrasocial"`
	Aportesindical                      *float64   `json:"aportesindical"`
	Importeretribucionesnohabituales    *float64   `json:"importeretribucionesnohabituales"`
	Importeretencionesgananciassufridas *float64   `json:"importeretencionesgananciassufridas"`
	Ajustes                             *float64   `json:"ajustes"`
	Importeconceptosexentos             *float64   `json:"importeconceptosexentos"`
	Sac                                 *float64   `json:"sac"`
	Importehorasextrasgravadas          *float64   `json:"importehorasextrasgravadas"`
	Importehorasextrasexentas           *float64   `json:"importehorasextrasexentas"`
	Materialdidactico                   *float64   `json:"materialdidactico"`
	Gastosmovilidad                     *float64   `json:"gastosmovilidad"`
}
