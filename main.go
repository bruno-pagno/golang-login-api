package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	fmt.Println("Starting server...")
	router := mux.NewRouter()
	router.HandleFunc("/", healthCheckHandler).Methods("GET")
	router.HandleFunc("/signin", signInHandler).Methods("POST")
	router.HandleFunc("/signup", signUpHandler).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
