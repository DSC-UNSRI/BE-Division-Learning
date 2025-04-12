package services

import (
    "fmt"
    "github.com/artichys/BE-Division-Learning/config"
    "github.com/artichys/BE-Division-Learning/utils"
)

func Authenticate() bool {
    fmt.Print("Masukkan Email: ")
    email := utils.ReadInput()
    fmt.Print("Masukkan Password: ")
    password := utils.ReadInput()

    if email == config.GetEnv("EMAIL") && password == config.GetEnv("PASSWORD") {
        fmt.Println("Login Berhasil!")
        return true
    }
    fmt.Println("Login Gagal!")
    return false
}
