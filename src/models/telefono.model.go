package models

//Telefono es el modelo de telefono en la base de datos
type Telefono struct {
	NoTelefono   string `json:"no_telefono,omitempty"`
	TipoTelefono int    `json:"tipo_telefono,omitempty"`
}
