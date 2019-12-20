package structSiradig

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Beneficiosiradig struct {
	structGormModel.GormModel
	Regionpatagonicameses        Mesesbeneficiosiradig `json:"regionpatagonicameses" gorm:"ForeignKey:Hijoid;association_foreignkey:ID;association_autoupdate:false"`
	Regionpatagonicamesesid      *int                  `json:"regionpatagonicamesesid" sql:"type:int REFERENCES Beneficiosiradigmeses(ID)"`
	Jubiladomeses                Mesesbeneficiosiradig `json:"jubiladomeses" gorm:"ForeignKey:Hijoid;association_foreignkey:ID;association_autoupdate:false"`
	Jubiladomesesid              *int                  `json:"jubiladomesesid" sql:"type:int REFERENCES Beneficiosiradigmeses(ID)"`
	Jubiladootrosingresosmeses   Mesesbeneficiosiradig `json:"jubiladootrosingresosmeses" gorm:"ForeignKey:Hijoid;association_foreignkey:ID;association_autoupdate:false"`
	Jubiladootrosingresosmesesid *int                  `json:"jubiladootrosingresosmesesid" sql:"type:int REFERENCES Beneficiosiradigmeses(ID)"`
	Tributobienes                bool                  `json:"tributobienes"`
	Tienemasbienes               bool                  `json:"tienemasbienes"`
}
