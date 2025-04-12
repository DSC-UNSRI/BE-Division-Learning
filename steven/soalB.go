package main

import (
	"fmt"
	"strconv"
)

func main() {
	var totalNumbers int
	fmt.Scan(&totalNumbers) 

	values := make([]int, totalNumbers)
	for i := 0; i < totalNumbers; i++ {
		fmt.Scan(&values[i])
	}

	for _, value := range values {
		numStr := strconv.Itoa(value) 
		reversedStr := ""

		for i := len(numStr) - 1; i >= 0; i-- {
			reversedStr += string(numStr[i])
		}

		reversedNum, _ := strconv.Atoi(reversedStr)
		fmt.Println(reversedNum)
	}
}
