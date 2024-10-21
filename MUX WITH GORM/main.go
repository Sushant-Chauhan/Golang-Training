package main

import (
	"BankingApp/models"
	"BankingApp/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection first
	models.InitDB()

	// Now, clear the database
	models.ClearDatabase()

	// Continue with routing and server setup
	router := mux.NewRouter()
	router = routes.SetupRouter(models.DB)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
