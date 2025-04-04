package main

import "fmt"

func sumOfNNaturalNumbers(N int32) int32 {

	var sum int32

	sum = (N * (N + 1)) / 2

	return sum
}
func main() {

	var N, answer int32

	N = 10
	fmt.Println("sum of the Natural number using the formula.")

	answer = sumOfNNaturalNumbers(N)
	fmt.Println("The sum of", N, "natural numbers is", answer)
}
