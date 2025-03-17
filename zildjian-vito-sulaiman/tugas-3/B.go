package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverseNumber(num int) int {
	strNum := strconv.Itoa(num)
	runes := []rune(strNum)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversed, _ := strconv.Atoi(strings.TrimLeft(string(runes), "0"))
	return reversed
}

func MainReverseNumber() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < N; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Fprintln(writer, reverseNumber(num))
	}
}
