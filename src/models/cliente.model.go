package models

//Cliente es el modelo de cliente en la base de datos
type Cliente struct {
	DPI       string `json:"dpi,omitempty"`
	Nombres   string `json:"nombres,omitempty"`
	Apellidos string `json:"apellidos,omitempty"`
	Edad      string `json:"edad,omitempty"`
}
