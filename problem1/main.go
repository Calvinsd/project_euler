package main

import "fmt"

func main() {
	ans1 := simpleSolution(1000)

	fmt.Println(ans1)
}

func simpleSolution(naturalNumber int) int {
	var sum int = 0
	for i := 1; i < naturalNumber; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

/*
	Formula for sum of an A.P like 2,4,6,8,10,12..... so on
	numberOfTerms * (startingTerm+lastTerm)/2

*/
