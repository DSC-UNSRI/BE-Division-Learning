package main

import (
	"bufio"
	"fmt"
	"os"
)

func sumDivisibleBy(x, N int) int {
	m := N / x
	return x * (m * (m + 1) / 2)
}

func sumMultiples(N int) int {
	return sumDivisibleBy(4, N) + sumDivisibleBy(7, N) - sumDivisibleBy(28, N)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)

	for i := 0; i < T; i++ {
		var N int
		fmt.Fscan(reader, &N)
		fmt.Fprintln(writer, sumMultiples(N))
	}
}
