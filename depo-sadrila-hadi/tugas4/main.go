package main

import (
	"bufio"
	"fmt"
	"os"
	"tugas4/controllers"
	"tugas4/models"
)

func main() {
    models.LoadEnv()
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Email: ")
    scanner.Scan()
    email := scanner.Text()
    fmt.Print("Password: ")
    scanner.Scan()
    password := scanner.Text()

    if !models.Authenticate(email, password) {
        fmt.Println("Email atau Password salah")
        return
    }

    data := models.NewData()
    controllers.ShowMenu(data)
}