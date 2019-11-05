package models

// Cuenta es el modelo de cuenta en la base de datos
type Cuenta struct {
	NoCuenta   int `json:"no_cuenta,omitempty"`
	Saldo      int `json:"saldo,omitempty"`
	Cliente    `json:"cliente,omitempty"`
	tipoCuenta `json:"tipo_cuenta,omitempty"`
}

type tipoCuenta struct {
	IDTipo      string `json:"id_tipo,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}
