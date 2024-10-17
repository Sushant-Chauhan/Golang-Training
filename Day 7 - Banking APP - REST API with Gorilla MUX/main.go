package main

import (
	"bankapp/controllers"
	middlewares "bankapp/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/create-super-admin", controllers.CreateSuperAdmin).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST") // Login to get JWT token: public route

	baseRouter := router.PathPrefix("/api/v1/bank-app").Subrouter()

	// subRouter - Banks & middleware
	subRouterForBanks := baseRouter.PathPrefix("/banks").Subrouter()
	subRouterForBanks.Use(middlewares.JWTAuthMiddleware) // Authentication middleware
	subRouterForBanks.Use(middlewares.AdminOnly)         // Admin-only middleware
	subRouterForBanks.HandleFunc("", controllers.GetAllBanks).Methods("GET")
	subRouterForBanks.HandleFunc("", controllers.CreateBank).Methods("POST")
	subRouterForBanks.HandleFunc("/{id}", controllers.GetBankByID).Methods("GET")
	subRouterForBanks.HandleFunc("/{id}", controllers.UpdateBankByID).Methods("PUT")
	subRouterForBanks.HandleFunc("/{id}", controllers.DeleteBankByID).Methods("DELETE")

	// subRouter - Users & middleware
	subRouterForUsers := baseRouter.PathPrefix("/users").Subrouter()
	subRouterForUsers.Use(middlewares.JWTAuthMiddleware)
	subRouterForUsers.Use(middlewares.AdminOnly)
	subRouterForUsers.HandleFunc("", controllers.GetAllUsers).Methods("GET")
	subRouterForUsers.HandleFunc("", controllers.CreateUser).Methods("POST")
	subRouterForUsers.HandleFunc("/{id}", controllers.GetUserByID).Methods("GET")
	subRouterForUsers.HandleFunc("/{id}", controllers.UpdateUserByID).Methods("PUT")
	subRouterForUsers.HandleFunc("/{id}", controllers.DeleteUserByID).Methods("DELETE")

	// Account routes
	router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/accounts/{id}", controllers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/accounts/{id}", controllers.DeleteAccount).Methods("DELETE")
	router.HandleFunc("/accounts/{id}/deposit", controllers.Deposit).Methods("POST")
	router.HandleFunc("/accounts/{id}/withdraw", controllers.Withdraw).Methods("POST")
	router.HandleFunc("/accounts/{fromID}/transfer/{toID}", controllers.Transfer).Methods("POST")

	// Start the server
	http.ListenAndServe(":8080", router)
}
