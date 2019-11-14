package models

// Transaccion es el modelo para las transacciones
type Transaccion struct {
	NoTransaccion int     `json:"no_transaccion,omitempty"`
	Fecha         string  `json:"fecha,omitempty"`
	NoCuenta      int     `json:"no_cuenta,omitempty"`
	TipoTran      string  `json:"tipo_tran,omitempty"`
	Monto         float64 `json:"monto,omitempty"`
}
