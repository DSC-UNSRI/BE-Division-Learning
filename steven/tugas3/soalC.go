package main

import "fmt"

func reverseNumber (num int) int {
	reverse := 0;
	for num > 0 {
		counter := num % 10;
		reverse = reverse * 10 + counter;
		num = num/10
	}
	return reverse;
}

func main() {
	var A, B, C, answer int;

	fmt.Scan(&A, &B);

	C = reverseNumber(A) + reverseNumber(B);

	answer = reverseNumber(C);
	fmt.Print(answer)
}