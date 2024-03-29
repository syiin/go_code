package routes

import (
	"controllers"
	"github.com/gorilla/mux"
	"net/http"
	"utils/auth"
)

func Handlers() *mux.Router {
	// 1. Create a new router
	// 2. Add the shared middleware - ie. setting content type and whatnot
	// 3. Add a sub-router to the new router
	// 4. Add the auth routes to the this subrouter with its own special middleware
	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	// These routes don't need auth
	r.HandleFunc("/", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	s.HandleFunc("/user", controllers.FetchUsers).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	// Transaction routes
	t := r.PathPrefix("/api").Subrouter()
	t.Use(auth.JwtVerify)
	t.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")
	t.HandleFunc("/transactions/{id}", controllers.GetTransaction).Methods("GET")
	t.HandleFunc("/transactions/{id}", controllers.UpdateTransaction).Methods("PUT")
	t.HandleFunc("/transactions/{id}", controllers.DeleteTransaction).Methods("DELETE")
	t.HandleFunc("/transactions", controllers.FetchTransactions).Methods("GET")
	t.HandleFunc("/transactions/search/", controllers.GetTransByField).Methods("GET")

	p := r.PathPrefix("/api").Subrouter()
	p.Use(auth.JwtVerify)
	p.HandleFunc("/propertyinfos", controllers.FetchPropertyInfos).Methods("GET")

	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
