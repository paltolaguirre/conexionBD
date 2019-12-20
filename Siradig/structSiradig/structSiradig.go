package structSiradig

import (
	"time"

	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Siradig struct {
	structGormModel.GormModel
	Legajo                                *structLegajo.Legajo                 `json:"legajo" gorm:"ForeignKey:Legajoid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Legajoid                              *int                                 `json:"legajoid" sql:"type:int REFERENCES Legajo(ID)" gorm:"not null"`
	Periodosiradig                        *time.Time                           `json:"periodosiradig" gorm:"not null"`
	Conyuge                               []Conyugesiradig                     `json:"conyuge" gorm:"ForeignKey:Conyugeid;association_foreignkey:ID;association_autoupdate:false"`
	Conyugeid                             *int                                 `json:"conyugeid" sql:"type:int REFERENCES Conyugesiradig(ID)"`
	Hijos                                 []Hijosiradig                        `json:"hijos" gorm:"ForeignKey:Hijoid;association_foreignkey:ID;association_autoupdate:false"`
	Hijoid                                *int                                 `json:"hijoid" sql:"type:int REFERENCES Hijosiradig(ID)"`
	Importegananciasotrosempleossiradig   *Importegananciasotrosempleossiradig `json:"legajo" gorm:"ForeignKey:Legajoid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Importegananciasotrosempleossiradigid *int                                 `json:"legajoid" sql:"type:int REFERENCES Importegananciasotrosempleossiradig(ID)" gorm:"not null"`
}
