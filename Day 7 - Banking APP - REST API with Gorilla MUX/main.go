package main

import (
	"bankingApp/controllers"
	"log"
	"net/http"
	"bankingApp/middlewares"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()


	router.HandleFunc("/login", controllers.LoginController).Methods(http.MethodPost) //authenticate

	adminRoutes := router.PathPrefix("/adminuser").Subrouter()
	adminRoutes.Use(middleware.AuthMiddleware)// authorize if admin
	adminRoutes.HandleFunc("/create", controllers.CreateCustomerController).Methods(http.MethodPost)
	adminRoutes.HandleFunc("/getcustomer", controllers.GetCustomerByIDController).Methods(http.MethodGet)
	adminRoutes.HandleFunc("/getall", controllers.GetAllCustomersController).Methods(http.MethodGet)	
	adminRoutes.HandleFunc("/update", controllers.UpdateCustomerController).Methods(http.MethodPut)
	adminRoutes.HandleFunc("/delete", controllers.DeleteCustomerController).Methods(http.MethodDelete)
	

	log.Println("Server is running on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
