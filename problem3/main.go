package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(simpleSolution(600851475143))
}

func simpleSolution(number int) int {

	for i := int(math.Sqrt(float64(number))); i > 0; i-- {
		fmt.Println("checking of", i)
		if isPrime(i) && number%i == 0 {
			return i
		}
	}

	return 1
}

func isPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true

}

// Interseting reads:
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
// Trial division
