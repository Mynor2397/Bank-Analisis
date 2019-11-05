package mssql

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

var (
	db   *sql.DB
	err  error
	once sync.Once
)

//Connect es la funcion para conectar a sqlServer
func Connect() (db *sql.DB) {
	server := `localhost`
	port := 1433
	user := "sa"
	password := "system"
	database := "Bank"

	once.Do(func() {
		connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
		db, err = sql.Open("sqlserver", connString)
		if err != nil {
			log.Println("Error en la conexion: ", err.Error())
		}
	})

	return db
}
