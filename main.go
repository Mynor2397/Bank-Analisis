package main

import (
	"log"
	"net/http"
	"os"

	r "github.com/Mynor2397/Bank-Analisis/src/routers"
)

func main() {
	//Miprincesa1009#
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := r.Router()
	log.Println("Listen And Serve on Port:", port)
	http.ListenAndServe(":"+port, router)
}
