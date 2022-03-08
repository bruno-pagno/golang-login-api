package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "accounts"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is up and running"))
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign in..."))
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign up..."))
}

func main() {
	fmt.Println("Connecting to the database")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	_, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Sucessfully connected to the database")
	}

	fmt.Println("Starting server...")
	router := mux.NewRouter()
	router.HandleFunc("/", healthCheckHandler).Methods("GET")
	router.HandleFunc("/signin", signInHandler).Methods("POST")
	router.HandleFunc("/signup", signUpHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
