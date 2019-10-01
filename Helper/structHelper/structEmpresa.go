package structHelper

type Empresa struct {
	ID                     int    `json:"id"`
	Nombre                 string `json:"nombre"`
	Codigo                 string `json:"codigo"`
	Descripcion            string `json:"descripcion"`
	Domicilio              string `json:"domicilio"`
	Localidad              string `json:"localidad"`
	Cuit                   string `json:"cuit"`
	Tipodeempresa          int    `json:"tipodeempresa"`
	Actividad              int    `json:"actividad"`
	Actividadnombre        string `json:"actividadnombre"`
	Zona                   int    `json:"zona"`
	Zonanombre             string `json:"zonanombre"`
	Obrasocial             int    `json:"obrasocial"`
	Artcontratada          int    `json:"artcontratada"`
	Domiciliodeexplotacion string `json:"domiciliodeexplotacion"`
	Reducevalor            int    `json:"reducevalor"`
}
