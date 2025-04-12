package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BacaInput(pesan string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(pesan)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input != "" {
			return input
		}
		fmt.Println("Input tidak boleh kosong, coba lagi.")
	}
}
