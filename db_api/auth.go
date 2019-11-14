package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	u "lens/utils"
	"net/http"
)

var JwtAuth = func (next http.Handler) http.Handler {
	return http.HandlerFunc(checkAuth)
}

var checkAuth = func (w http.ResponseWriter, r *http.Request) {
	notAuth := []string{"/api/user/new", "/api/user/login"}
	requestPath := r.URL.Path //current request path

	//check if request needs auth
	for _,value := range notAuth {
		if value == requestPath{
			next.ServeHTTP(w, r)
			return
		}
	}

	res := make(map[string] interface{})
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" { //token is missing
		res = u.Message(false, "Missing auth token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, res)
		return
	}

	splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		res = u.Message(false, "Invalid/Malformed auth token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}

	tokenPart := splitted[1] //Grab the token part
	tk := &models.Token{}
	token, err := jwt.ParseWithClaims(tokenPart, 
									tk, 
									func(token *jwt.Token) (interface{}, error) {
										return []byte(os.Getenv("TOKEN_PASSWORD")), nil})
	 if err != nil { //malformed token, returns with 403
		 res = u.Message(false, "Malformed auth token")
		 w.WriteHeader(http.StatusForbidden)
		 w.Header().Add("Content-Type", "application/json")
		 u.Respond(w, res)
		 return
	 }

	 if !token.Valid { //token is invalid
		 res = u.Message(false, "Token is invalid")
		 w.WriteHeader(http.StatusForbidden)
		 w.Header().Add("Content-Type", "application/json")
		 u.Respond(w, res)
		 return
	 }

	 fmt.Sprintf("User %", tk.Username)
	 ctx := context.WithValue(r.Context(), "user", tk.UserId)
	 r = r.WithContext(ctx)
	 next.ServeHTTP(w, r) //proceed with middleware chain
}



