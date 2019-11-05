package models

// Transacciones es el modelo para las transacciones
type Transacciones struct {
	NoTransaccion   int    `json:"no_transaccion,omitempty"`
	Fecha           string `json:"fecha,omitempty"`
	NoCuenta        int    `json:"no_cuenta,omitempty"`
	tipoTransaccion `json:"tipo_transaccion,omitempty"`
}

type tipoTransaccion struct {
	IDTipo      string `json:"id_tipo,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}
