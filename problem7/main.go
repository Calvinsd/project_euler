package main

import (
	"fmt"
	"math"
)

func main() {
	var n = 10001
	var primeCount = 0
	var primeNumber = 2

	for primeCount != n {
		if isPrime(primeNumber) {
			primeCount++
			if primeCount != n {
				primeNumber++
			}
			continue
		}

		primeNumber++
	}

	fmt.Println("Result :", primeNumber)
}

func isPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true

}
