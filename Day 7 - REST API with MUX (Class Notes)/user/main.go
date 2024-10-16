package main

//objective CRUD on USER with Credentials
import (
	"encoding/json"
	"net/http"
	"time"
	"user/components/user/controller"
	"user/components/user/service"
	"user/middleware"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", login).Methods(http.MethodPost)
	// router.PathPrefix("/api/v1")
	subRouterForMiddleware1 := router.NewRoute().Subrouter()
	// subRouterForMiddleware2 := router.NewRoute().Subrouter()
	// subRouterForMiddleware1.PathPrefix("/user")
	// subRouterForMiddleware2.PathPrefix("/user")
	subRouterForMiddleware1.HandleFunc("/api/v1/user/", controller.GetUserByID).Methods(http.MethodGet) //1
	// subRouterForMiddleware1.HandleFunc("/{id}", controller.GetUserByID).Methods(http.MethodGet) //2
	// subRouterForMiddleware2.HandleFunc("/{id}", controller.UpdateUserByID).Methods(http.MethodPut) //3
	// subRouterForMiddleware2.HandleFunc("/", controller.CreateUser).Methods(http.MethodPost)        //4 "/api/v1/user"
	subRouterForMiddleware1.Use(middleware.VerifyAdmin)
	// subRouterForMiddleware2.Use(middleware.Middleware2)
	router.Use(middleware.Middleware0)
	http.ListenAndServe(":4000", router)
}
func login(w http.ResponseWriter, r *http.Request) {
	var user service.User
	json.NewDecoder(r.Body).Decode(&user)
	//validation
	//bussness Logic
	//creating a payload
	claim := middleware.NewClaims(1, user.Username, true, time.Now().Add(time.Minute*3))
	token := claim.Signing()
	w.Header().Set("Authorization", token)
	json.NewEncoder(w).Encode(claim)
	//header Authorization
}

//middleware
