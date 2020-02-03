package structFormula

import (
	"encoding/json"

	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Formula struct {
	structGormModel.GormModel
	Key   string `json:"key"`
	Value string `json:"value" sql:"type:json"`
}

type FormulaPrime struct {
	Signature Signature `json:"signature"`
	Invoke    Invoke    `json:"invoke"`
}

type Signature struct {
	Function string          `json:"function"`
	Result   string          `json:"result"`
	Param    json.RawMessage `json:"param"`
}

type Invoke struct {
	Function string          `json:"function"`
	Arg      json.RawMessage `json:"arg"`
}

type FormulaWrapper struct {
	structGormModel.GormModel
	FormulaPrime
}
