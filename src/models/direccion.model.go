package models

// Direccion es el model de Direcciones en la base de datos
type Direccion struct {
	CodDireccion  string `json:"cod_direccion,omitempty"`
	Descripcion   string `json:"descripcion,omitempty"`
	TipoDireccion string `json:"tipo_direccion,omitempty"`
	Cliente       `json:"cliente,omitempty"`
}
