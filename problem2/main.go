package main

import "fmt"

func main() {

	ans := simpleSolution(4000000)

	fmt.Println(ans)

}

func simpleSolution(limit int) int {
	firstNumber := 1
	secondNumber := 2
	var nextNumber int

	evenSum := 2

	if limit < 2 {
		return 0
	}

	if limit == 2 {
		return secondNumber
	}

	for {
		nextNumber = firstNumber + secondNumber
		firstNumber = secondNumber
		secondNumber = nextNumber

		if nextNumber > limit {
			break
		}

		if nextNumber%2 == 0 {
			evenSum += nextNumber
		}

	}

	return evenSum
}
