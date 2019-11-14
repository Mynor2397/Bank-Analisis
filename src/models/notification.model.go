package models

// Notification es para reportar el mensaje a Adelsa
type Notification struct {
	NoCuenta int     `json:"no_cuenta,omitempty"`
	DPI      int     `json:"dpi,omitempty"`
	Monto    float64 `json:"monto,omitempty"`
	Razon    string  `json:"razon,omitempty"`
	Status   bool    `json:"status"`
}
