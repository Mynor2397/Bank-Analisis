package models

//Cliente es el modelo de cliente en la base de datos
type Cliente struct {
	DPI       string `json:"dpi,omitempty"`
	Nombres   string `json:"nombres,omitempty"`
	Apellidos string `json:"apellidos,omitempty"`
	FechaNac  string `json:"fecha_nac,omitempty"`
	Telefono  `json:"telefono,omitempty"`
	Direccion `json:"direccion,omitempty"`
	Cuenta    `json:"cuenta,omitempty"`
}
