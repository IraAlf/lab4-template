package main

import (
	"log"
	"net/http"

	"lab2/src/gateway-service/internal/handlers"
)

func main() {
	port := "8080"
	r := handlers.Router()
	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
