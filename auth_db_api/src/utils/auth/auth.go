package auth

import (
	"models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//Exception struct
type Exception models.Exception

// JwtVerify Middleware function
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//1. Get the token string
		//2. Create a key function to get the key
		//3. Get the Claim used to sign the token string
		//4. Give all 3 to the Parser to validate
		var header = r.Header.Get("x-access-token") 
		header = strings.TrimSpace(header)
		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		tk := &models.Token{}
		out, _ := json.Marshal(tk) 
		fmt.Println("\nBefore tk was populated: ")
		fmt.Println(string(out))

		keyFunc := func(token *jwt.Token) (interface{}, error) {
					return []byte("secret"), nil }
		reconstructed, err := jwt.ParseWithClaims(header, tk, keyFunc)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		//just to visualise the return struct
		out, _ = json.Marshal(reconstructed) 
		fmt.Println("\nFrom JwtVerify(): ")
		fmt.Println(string(out))
		//notice that tk has been populated by ParseWithClaims()
		out, _ = json.Marshal(tk) 
		fmt.Println("\nFrom tk populated with: ")
		fmt.Println(string(out))

		//a context is an interface, an interface is a description of method 
		//signatures an object can have
		//contexts define boundaries in the program (eg. stopping, channels)
		//"user" is the key - WithValue() takes a context and creates a derivative
		//context associated with the value (ie. tk)
		//this is a way to pass on information - this information is local to
		//the request and is then deleted once the request is done
		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
