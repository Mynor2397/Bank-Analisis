package mssql

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/denisenkom/go-mssqldb" //Es el controlador para sql server
)

var (
	db   *sql.DB
	err  error
	once sync.Once
)

//Connect es la funcion para conectar a sqlServer
func Connect() (db *sql.DB) {

	// server := string("MAICK\\mssqlserverd")
	// port := 1433
	// user := "sa"
	// password := "root"
	// database := "Bank"


	server := "adelsabank.database.windows.net"
	port := 1433
	user := "adelsa"
	password := "Proyectofinal1"
	database := "BancoAdelsa"


	once.Do(func() {
		connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
		db, err = sql.Open("sqlserver", connString)
		if err != nil {
			log.Println("Error en la conexion: ", err.Error())
		}
	})

	return db
}
