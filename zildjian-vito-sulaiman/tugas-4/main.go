package main

import (
	"fmt"

	"tugas-4/controllers"
)

func main() {

	user, err := controllers.LoadEnv()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !controllers.Authenticate(user) {
		fmt.Println("Invalid email or password!")
		return
	}

	controllers.StartDashboard(user)
}
