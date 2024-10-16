// sets up a simple HTTP server using Gorilla Mux router to handle POST requests to the /api/v1/mux/{id} endpoint. Decodes a JSON body into a credential struct and prints headers, route parameters, and query parameters.

package main

import (
	"encoding/json" // handle JSON encoding and decoding.
	"fmt"
	"net/http" //Provides HTTP server functionality

	"github.com/gorilla/mux" //for routing
)

// here struct is primarily used to store the body data from incoming requests.
type credential struct {
	Username string `json:"username"` // JSON tags, which helps in decoding the JSON request body. serialized and deserialized from JSON using the key "username".
	Password string `json:"password"`
}

func main() {
	router := mux.NewRouter()                                      //create a new Gorilla Mux router using mux.NewRouter()    ( allows defining - 	URL routes and associating them with handler functions)
	router.HandleFunc("/api/v1/mux/{id}", muxDemo).Methods("POST") //POST requests - it trigger  muxDemo handler function. {id}- route parameter
	fmt.Println("Server is running on port 4000...")
	http.ListenAndServe(":4000", router) //starts the HTTP server on port 4000, serving requests using the router.
}

// muxDemo function - core of the API logic. It handles request and response
func muxDemo(w http.ResponseWriter, r *http.Request) { // body , headers, routeParams, queryparam

	//1. Reading the Body
	var cred = &credential{}                    //request body (which is expected to be JSON) is decoded into the cred struct, storing the credentials (Username, Password) sent by the client.
	err := json.NewDecoder(r.Body).Decode(cred) //if the body is not valid JSON
	if err != nil {
		panic(err.Error())
	}

	// 2. Reading Headers
	headers := r.Header //Reading Headers - contains meta-information about the request, like authentication tokens, content type, etc.

	// 3. Extracting Route Parameters
	routeParams := mux.Vars(r)

	// 4. Reading Query Parameters - Appears after the ? in the URL.
	queryParams := r.URL.Query()
	fmt.Println("Received request with query parameters:", queryParams)

	// 5. Printing Information
	fmt.Println("headers>>>", headers)
	fmt.Println("routeParams>>>", routeParams)
	fmt.Println("queryParams>>>", queryParams)

	// 6. Setting Response Headers and Status
	w.Header().Set("yash", "shah") //response header yash is set to shah -response header yash is set to shah
	w.WriteHeader(http.StatusOK)   //The response status is set to 200 OK.

	// 7. Sending the Response
	json.NewEncoder(w).Encode(cred) //cred struct is encoded back into JSON and sent as the response body (cred struct - 	which contains the Username and Password sent in the request)

}

/*
Key Concepts in the Code:
Routing: Handled by Gorilla Mux, which allows defining routes with dynamic URL parameters.
Request Handling: The handler function reads the request body, headers, and parameters, processes them, and sends an appropriate response.
JSON Parsing: The request body is parsed into a Go struct using - json.NewDecoder().Decode(), and responses are sent using - json.NewEncoder().Encode().
Headers, Query Params, Route Params: These are different ways of sending data in HTTP requests:
Headers: Meta-information about the request.
Query Parameters: Additional data sent in the URL after ?.
Route Parameters: Dynamic parts of the URL, like {id}.
*/

//CRUD
//Authentication in MUX
