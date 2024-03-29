package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(e)
	port := os.Getenv("PORT")

	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
