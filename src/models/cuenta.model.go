package models

// Cuenta es el modelo de cuenta en la base de datos
type Cuenta struct {
	NoCuenta    int     `json:"no_cuenta,omitempty"`
	Saldo       float32 `json:"saldo,omitempty"`
	IDTipo      string  `json:"id_tipo,omitempty"`
	Descripcion string  `json:"descripcion,omitempty"`
	Cantidad    int     `json:"cantidad,omitempty"`
}
