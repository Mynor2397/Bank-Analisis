package models

// Direccion es el model de Direcciones en la base de datos
type Direccion struct {
	DescripcionD  string `json:"descripcion_d,omitempty"`
	TipoDireccion string `json:"tipo_direccion,omitempty"`
}
