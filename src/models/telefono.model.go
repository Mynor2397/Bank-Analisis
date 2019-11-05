package models

//Telefono es el modelo de telefono en la base de datos
type Telefono struct {
	NoTelefono   string `json:"no_telefono,omitempty"`
	TipoTelefono string `json:"tipo_telefono,omitempty"`
	Cliente      `json:"cliente,omitempty"`
}
