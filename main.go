package main

import (
	"log"
	"net/http"
	"oob/backend/handler"
)

func main() {
	// Création du routeur
	r := handler.NewRouter()

	// Démarrage du serveur
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
