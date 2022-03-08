package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
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
	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Connecting to the database")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

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
