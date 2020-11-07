package main

import "fmt"

/*
The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

func prob6() int {

	var (
		sumOfSquare int
		sumOfNum    int
	)

	for i := 1; i <= 100; i++ {
		sumOfSquare += i * i
		sumOfNum += i
	}

	return sumOfNum*sumOfNum - sumOfSquare
}

func main() {
	fmt.Println(prob6())
}
