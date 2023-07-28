package main

import "fmt"

func main() {
	var startingPoint = 2520 / 20 // as 2520 is the smallest number divisible  by upto all first 10 digits
	var smallestPositiveNumber int
	var iterations int

	for true {
		// iterate over multiple of 20s as first it should be divisible by 20
		number := startingPoint * 20
		if number%20 == 0 && number%19 == 0 && number%18 == 0 &&
			number%17 == 0 && number%16 == 0 && number%15 == 0 &&
			number%14 == 0 && number%13 == 0 && number%12 == 0 && number%11 == 0 {
			smallestPositiveNumber = number
			break
		}

		startingPoint += 1
		iterations += 1
	}

	fmt.Printf("Smallest positive number after iterstions %d is %d\n", iterations, smallestPositiveNumber)
}
