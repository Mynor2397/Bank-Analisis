package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	help "github.com/Mynor2397/Bank-Analisis/src/helpers"
	"github.com/Mynor2397/Bank-Analisis/src/models"
)

// Deposito es para el ingreso del dinero a la cuenta bancaria
func Deposito(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	defer r.Body.Close()
	deposit := models.Transaccion{}

	json.NewDecoder(r.Body).Decode(&deposit)

	tsql := fmt.Sprintf("exec SP_DEPOSITO '%d', '%s', %f", deposit.NoCuenta, deposit.TipoTran, deposit.Monto)
	Query, err := db.Query(tsql)

	if err == nil {
		notification := models.Notification{
			NoCuenta: deposit.NoCuenta,
			Monto:    deposit.Monto,
			Razon:    "Transaccion realizada exitosamente",
			Status:   true,
		}

		jsonresult, _ := json.Marshal(notification)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonresult)
		return
	}

	if err.Error() == help.ErrorCuentaNotFound {
		notification := models.Notification{
			NoCuenta: deposit.NoCuenta,
			Monto:    deposit.Monto,
			Razon:    "El numero de cuenta proporcionado no es válido",
			Status:   false,
		}

		jsonresult, _ := json.Marshal(notification)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonresult)
		return
	}

	if err != nil {
		log.Println("+++ Error no controlado: ", err.Error(), "+++")
		return
	}

	defer Query.Close()
}

// Debito es para el debito automático en la cuenta del cliente
func Debito(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	defer r.Body.Close()
	debito := models.Transaccion{}

	json.NewDecoder(r.Body).Decode(&debito)

	tsql := fmt.Sprintf("EXEC SP_DEBITO %d, %f", debito.NoCuenta, debito.Monto)
	Query, err := db.Query(tsql)

	if err == nil {
		notification := models.Notification{
			NoCuenta: debito.NoCuenta,
			Monto:    debito.Monto,
			Razon:    "El estado se encuentra correcto",
			Status:   true,
		}

		jsonresult, _ := json.Marshal(notification)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonresult)
		return
	}

	if err.Error() == help.ErrorCuentaNotFound {
		notification := models.Notification{
			NoCuenta: debito.NoCuenta,
			Monto:    debito.Monto,
			Razon:    "El numero de cuenta proporcionado no es válido",
			Status:   false,
		}

		jsonresult, _ := json.Marshal(notification)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonresult)
		return
	}

	if err.Error() == help.ErrorSaldoInsuficiente {
		notification := models.Notification{
			NoCuenta: debito.NoCuenta,
			Monto:    debito.Monto,
			Razon:    "El saldo es insuficiente",
			Status:   false,
		}

		jsonresult, _ := json.Marshal(notification)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonresult)
		return
	}

	if err != nil {
		log.Println("+++ Error no controlado: ", err.Error(), "+++")
		return
	}

	defer Query.Close()

}
