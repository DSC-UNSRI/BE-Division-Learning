package main

import "fmt"

func main() {
	var email string;
	var password string;

	fmt.Println("Login Dashboard Iftar GDGoC")
	fmt.Print("EMAIL: ")
	fmt.Scan(&email)
	fmt.Print("PASSWORD: ")
	fmt.Scan(&password)

	fmt.Println(email)
	fmt.Println(password)
}