// folder :
// components > user> controller > userController.go 

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/components/user/service"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserByID Controller Called")
	// Extract user ID from the URL or request parameters
	userID := 1 // Replace with actual extraction logic from the URL
	user, err := service.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user service.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newUser, err := service.NewUser(user.Username, user.Password, user.Name, user.Age, user.IsAdmin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newUser)
}
