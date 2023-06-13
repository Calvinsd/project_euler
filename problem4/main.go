package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(1221))
}

func isPalindrome(number int) bool {
	var reverseNumber int
	n := number

	for n > 0 {
		reverseNumber = (reverseNumber * 10) + (n % 10)
		n = n / 10
	}

	if reverseNumber == number {
		return true
	}

	return false
}
