package structSiradig

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Mesesbeneficiosiradig struct {
	structGormModel.GormModel
	Tipobeneficio   Tipobeneficiosiradig `json:"tipobeneficio" gorm:"ForeignKey:Tipobeneficioid;association_foreignkey:ID;association_autoupdate:false"`
	Tipobeneficioid *int                 `json:"tipobeneficioid" sql:"type:int REFERENCES Tipobeneficiosiradig(ID)"`
	Enero           bool                 `json:"enero"`
	Febrero         bool                 `json:"febrero"`
	Marzo           bool                 `json:"marzo"`
	Abril           bool                 `json:"abril"`
	Mayo            bool                 `json:"mayo"`
	Junio           bool                 `json:"junio"`
	Julio           bool                 `json:"julio"`
	Agosto          bool                 `json:"agosto"`
	Septiembre      bool                 `json:"septiembre"`
	Octubre         bool                 `json:"octubre"`
	Noviembre       bool                 `json:"noviembre"`
	Diciembre       bool                 `json:"diciembre"`
}
