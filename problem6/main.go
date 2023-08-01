package main

import (
	"fmt"
	"math"
)

/*
Formula for sum of the squares of the first n natural numbers  is: n(n+1)(2n+1)/6

Formula for sum of the first n natural numbers: n(n+1)/2

Sum of the first  n odd numbers is n^2
*/
func main() {
	var n = 100

	sumOfSquares := (n * (n + 1) * (2*n + 1)) / 6

	squareOfSum := int(math.Pow(float64((n*(n+1))/2), 2))

	fmt.Println("Result :", squareOfSum-sumOfSquares)
}
