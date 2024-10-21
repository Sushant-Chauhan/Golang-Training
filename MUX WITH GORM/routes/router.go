package routes

import (
	"BankingApp/controllers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// Route for creating the Super Admin
	router.HandleFunc("/create-super-admin", controllers.CreateSuperAdminController(db)).Methods("POST")

	//// Route for logging in and generating JWT token
	// router.HandleFunc("/login", controllers.LoginController(db)).Methods("POST")

	//// Secured routes
	userRouter := router.PathPrefix("/users").Subrouter()
	// userRouter.Use(middlewares.TokenAuthMiddleware) // Require authentication for user routes
	userRouter.HandleFunc("/", controllers.CreateUserController(db)).Methods("POST")
	userRouter.HandleFunc("/{id}", controllers.GetCustomerByIDController(db)).Methods("GET")
	userRouter.HandleFunc("/", controllers.GetAllCustomersController(db)).Methods("GET")
	userRouter.HandleFunc("/{id}", controllers.UpdateCustomerController(db)).Methods("PUT")
	userRouter.HandleFunc("/{id}", controllers.DeleteCustomerController(db)).Methods("DELETE")

	// Account routes for basic CRUD operations
	router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts", controllers.GetAllAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", controllers.GetAccountByID).Methods("GET")
	router.HandleFunc("/accounts/{id}", controllers.UpdateAccountByID).Methods("PUT")
	router.HandleFunc("/accounts/{id}", controllers.DeleteAccountByID).Methods("DELETE")

	// Routes for handling transactions within an account's passbook
	router.HandleFunc("/accounts/{id}/passbook", controllers.GetPassbook).Methods("GET")

	// Ledger routes
	router.HandleFunc("/ledger", controllers.CreateLedgerEntry).Methods("POST")
	router.HandleFunc("/ledger/{bankID}", controllers.GetLedgerEntries).Methods("GET")

	// Routes for account transactions
	router.HandleFunc("/accounts/{id}/withdraw", controllers.WithdrawFromAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}/deposit", controllers.DepositToAccount).Methods("POST")
	router.HandleFunc("/accounts/transfer", controllers.TransferBetweenAccounts).Methods("POST")

	// Bank routes
	router.HandleFunc("/banks", controllers.CreateBank).Methods("POST")
	router.HandleFunc("/banks", controllers.GetAllBanks).Methods("GET")
	router.HandleFunc("/banks/{id}", controllers.GetBankByID).Methods("GET")
	router.HandleFunc("/banks/{id}", controllers.UpdateBankByID).Methods("PUT")
	router.HandleFunc("/banks/{id}", controllers.DeleteBankByID).Methods("DELETE")

	return router
}
