package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumMultiples(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%4 == 0 || i%7 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Baca jumlah pertanyaan
	input, _ := reader.ReadString('\n')
	T, _ := strconv.Atoi(input)

	// Proses setiap nilai N
	for i := 0; i < T; i++ {
		input, _ := reader.ReadString('\n')
		N, _ := strconv.Atoi(input)
		result := sumMultiples(N)
		fmt.Fprintln(writer, result)
	}
}
