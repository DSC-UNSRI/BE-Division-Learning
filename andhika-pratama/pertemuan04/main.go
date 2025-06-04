package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "pertemuan04/controllers"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

    email := os.Getenv("EMAIL")
    password := os.Getenv("PASSWORD")

    fmt.Print("Enter email: ")
    inputEmail := readInput()
    fmt.Print("Enter password: ")
    inputPassword := readInput()

    if inputEmail != email || inputPassword != password {
        fmt.Println("Invalid credentials")
        return
    }

    fmt.Println("Login successful!")
    controllers.Dashboard()
}

func readInput() string {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}
