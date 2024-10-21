package main

//objective CRUD on USER with Credentials
import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"user/components/user/controller"
	"user/components/user/service"
	"user/middleware"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", login).Methods(http.MethodPost) //POST URL  ---  http://localhost:4000/login
	// router.PathPrefix("/api/v1")
	subRouterForMiddleware1 := router.NewRoute().Subrouter()
	subRouterForMiddleware1.HandleFunc("/api/v1/user/", controller.GetUserByID).Methods(http.MethodGet) //1   GET URL ---  http://localhost:4000/api/v1/user/
	subRouterForMiddleware1.Use(middleware.Authentication)
	// router.Use(middleware.Middleware0)
	fmt.Println("Server started ...")
	http.ListenAndServe(":4000", router)
}

func login(w http.ResponseWriter, r *http.Request) {
	var user service.User
	json.NewDecoder(r.Body).Decode(&user)
	//validation
	//bussness Logic
	//creating a payload
	claim := middleware.NewClaims(1, user.Username, true, time.Now().Add(time.Minute*3000))
	token := claim.Signing()
	w.Header().Set("Authorization", token)
	json.NewEncoder(w).Encode(claim)
	//header Authorization
}
