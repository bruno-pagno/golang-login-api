package main

import (
	"fmt"
)

func signIn() {
	fmt.Println("Signing in...")
}

func signUp(email string, password string) {
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println("Signing up...")
}

func main() {
	signIn()
	signUp("brunopagno@usp.br", "batata1234")
}
