package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/components/user/service"
)

// func GetAllUsers(w http.ResponseWriter, r *http.Request) {
// 	//GET
// 	//send recorde based on limit/pageSize and offset/pageNo
// 	// 5 pageSize ; 3 pageNo //pagination
// 	//limit and offset ??
// 	//default value; queryParams??
// 	//validation limit and offset

// 	a, totalCount := service.GetAllUsers(5, 3)

// 	//body,stuscode
// 	//body {
// 	// totalCount:,
// 	// data:a[]
// 	// }
// 	//body : data:a[]
// 	//headers:   x-total-count:totacCOunt // data exchange
// }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	// // data Username
// 	// Password
// 	// Name
// 	// Age
// 	// IsAdmin
// 	//body json {Username
// 	// Password
// 	// Name
// 	// Age
// 	// IsAdmin}

// 	//validation
// 	user, err := service.NewUser()
// 	if err != nil {
// 		//statuscode badReq
// 		// return response with body as {errorMessage: err.Error()}
// 		json.NewEncoder(w).Encode(err)
// 		// http.Error(w,)
// 		return
// 	}

// 	//res body -->user

// }

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserByID Controller Called")
	service.GetUserByID()
	json.NewEncoder(w).Encode(&service.User{})
}

// func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
// 	//validations

// 	//service call
// 	user:=&service.User{}
// 	service.UpdateUserByID()

// }
