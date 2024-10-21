// folder :
// middleware > middlware.go

package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func Middleware0(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware0 Called")
		next.ServeHTTP(w, r)
	})
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authentication Header Started >>>>>")
		//get
		tokenString, ok := r.Header["Authorization"]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errors.New("Token not found").Error())
		}
		claim := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString[0], claim, func(t *jwt.Token) (interface{}, error) {
			return secrateKey, nil
		})
		fmt.Println("claims generated >>>>>", claim)
		fmt.Print("token parser >>>>>>>>", token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errors.New("Token Invalid").Error())
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errors.New("Token Expired").Error())
			return
		}

		// request --> claims
		// _ = context.WithValue(r.Context(), "claims", claim)
		fmt.Print("Auth Middleware Ended >>>>>>>>>>>")
		next.ServeHTTP(w, r)
	})
}

// func VerifyAdmin(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(" ")
// 		//get
// 		claim, ok := r.Context().Value("claims").(*Claims)
// 		if !ok || !claim.IsAdmin {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			json.NewEncoder(w).Encode(errors.New("Unauthorized access").Error())
// 			return
// 		}
// 		if !claim.IsAdmin {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			json.NewEncoder(w).Encode(errors.New("IsAdmin False").Error())
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

func Middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware2 Called")
		next.ServeHTTP(w, r)
	})
}
