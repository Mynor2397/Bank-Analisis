package routers

import (
	con "github.com/Mynor2397/Bank-Analisis/src/controllers"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

// Router devuelve un enrutador por cada petición
func Router() *mux.Router {
	router.HandleFunc("/cliente", con.CreateClient).Methods("POST")
	router.HandleFunc("/deposito", con.Deposito).Methods("POST")
	router.HandleFunc("/debito", con.Debito).Methods("POST")
	router.HandleFunc("/cliente/{nocuenta}", con.OnlyAcount).Methods("GET")
	return router
}
