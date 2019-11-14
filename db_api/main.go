package main

import (
	"db_api/handlers"
	"database/sql"
	_"github.com/lib/pq" //'_' character means we do not interact directly with the postgres driver
	"os"
	"fmt"
	"log"
	"net/http"
)

var db *sql.DB
const  (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main(){
	initDb()
	defer db.Close()

	http.HandleFunc("/api/index", handlers.IndexHandler(db))
	http.HandleFunc("/api/repo/", handlers.TransactionsHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func initDb() {
    config := dbConfig()
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        config[dbhost], config[dbport],
        config[dbuser], config[dbpass], config[dbname])

    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
	        panic(err)
    }
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string { //a map is basically a hash
	conf := make(map[string]string) //make allocates and initialise a map value that points
	host, ok := os.LookupEnv(dbhost) //looks for the environment var
	if !ok {
		panic("DBHOST env var required but not set")
	}
	
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT env var required but not set")
	}

	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER env var required but not set")
	}
	
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS env var required but not set")
	}
	
	name, ok := os.LookupEnv(dbname);
	if !ok {
		panic("DBNAME env var required but not set")
	}

	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}