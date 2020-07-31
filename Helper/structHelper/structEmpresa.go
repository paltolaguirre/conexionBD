package structHelper

type Empresa struct {
	ID                     int    `json:"ID"`
	Nombre                 string `json:"nombre"`
	Codigo                 string `json:"codigo"`
	Descripcion            string `json:"descripcion"`
	Domicilio              string `json:"domicilio"`
	Localidad              string `json:"localidad"`
	Cuit                   string `json:"cuit"`
	Actividad              string `json:"actividad"`
	Actividadnombre        string `json:"actividadnombre"`
	Tipodeempresa          int    `json:"tipodeempresa"`
	Zonaid                 int    `json:"zonaid"`
	Zona                   string `json:"zona"`
	Zonanombre             string `json:"zonanombre"`
	Obrasocial             int    `json:"obrasocial"`
	Artcontratada          int    `json:"artcontratada"`
	Domiciliodeexplotacion string `json:"domiciliodeexplotacion"`
	Reducevalor            int    `json:"reducevalor"`
	Logo                   string `json:"logo"`
}
