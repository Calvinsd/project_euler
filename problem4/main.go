package main

import "fmt"

func main() {
	var largestPalindrome int

	// brute force
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			if isPalindrome(i*j) && (i*j) > largestPalindrome {
				fmt.Printf(" %d * %d = %d \n", i, j, i*j)
				largestPalindrome = i * j
			}
		}
	}

	fmt.Println("Largest palindrome", largestPalindrome)
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
