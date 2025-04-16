package main

import "fmt"


func reverseNumber(num int) int {

	temp := num
	reverse := 0

	for temp > 0 {
		reverse = reverse*10 + temp%10
		temp /= 10
	}
	return reverse
 }


 
func main() {
	
	var a, b int

	fmt.Scan(&a, &b)

	reversedA := reverseNumber(a)
	reversedB := reverseNumber(b)

	hasil := reversedA + reversedB

	hasilReversed := reverseNumber(hasil)

	fmt.Println(hasilReversed)

}
