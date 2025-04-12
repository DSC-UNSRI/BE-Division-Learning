package utils

import "fmt"

func TanganiError(err error, pesan string) {
	if err != nil {
		fmt.Println("Error:", pesan)
	}
}
