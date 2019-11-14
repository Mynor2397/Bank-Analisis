package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	help "github.com/Mynor2397/Bank-Analisis/src/helpers"
	mod "github.com/Mynor2397/Bank-Analisis/src/models"
	"github.com/Mynor2397/Bank-Analisis/src/services/mssql"
	"github.com/gorilla/mux"
)

var (
	db = mssql.Connect()
)

// CreateClient registra un nuevo cliente en la base de datos
func CreateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	defer r.Body.Close()
	cliente := mod.Cliente{}

	json.NewDecoder(r.Body).Decode(&cliente)

	var rn int

	for {
		rn = rand.New(rand.NewSource(time.Now().UnixNano())).Int() % 100000000
		if rn > 10000000 {
			break
		}
	}

	ncuenta := fmt.Sprintf("%d", rn)
	cliente.NoCuenta, _ = strconv.Atoi(ncuenta)

	log.Println(cliente)
	tsql := fmt.Sprintf(`exec SP_CREARCLIENTE '%s', '%s', '%s', '%s', %s, %d, '%s', '%s', '%d', '%s'`,
		cliente.DPI, cliente.Nombres, cliente.Apellidos, cliente.FechaNac, cliente.Telefono.NoTelefono, cliente.Telefono.TipoTelefono, cliente.Direccion.DescripcionD, cliente.Direccion.TipoDireccion, cliente.NoCuenta, cliente.Cuenta.IDTipo)

	query, err := db.Query(tsql)

	if err == nil {
		jsonresult, _ := json.Marshal(cliente)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonresult)
		return
	}

	if err.Error() == help.AccountExist {
		notification := mod.Notification{
			DPI:    cliente.DPI,
			Razon:  "La cuenta ya existe",
			Status: false,
		}

		jsonresult, _ := json.Marshal(notification)
		w.WriteHeader(http.StatusConflict)
		w.Write(jsonresult)
		return
	}

	if err != nil {
		log.Println("Error en el SP: ", err.Error())
		http.Error(w, help.ErrInsert, http.StatusConflict)
		return
	}

	defer query.Close()

}

//OnlyAcount es para traer un solo clientes
func OnlyAcount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	nocuenta := params["nocuenta"]

	cuenta := mod.Cliente{}

	tsql := fmt.Sprintf("SELECT * FROM onlyclient('%s', '%s')", nocuenta, nocuenta)

	Query, err := db.Query(tsql)
	if err != nil {
		log.Println("Error en la function: ", err.Error())
		http.Error(w, help.ErrInsert, http.StatusBadRequest)
		return
	}

	for Query.Next() {
		err := Query.Scan(&cuenta.NoCuenta, &cuenta.Nombres, &cuenta.Apellidos, &cuenta.Saldo, &cuenta.DPI, &cuenta.DescripcionD, &cuenta.NoTelefono)
		if err != nil {
			log.Println("Error en la consulta en el scanner de la funcion:", err.Error())
		}
	}

	if cuenta.Nombres == "" {
		http.Error(w, help.ErrNoData, http.StatusNoContent)
	} else {
		output, _ := json.Marshal(cuenta)
		fmt.Fprintf(w, string(output))
	}

	defer Query.Close()
}
