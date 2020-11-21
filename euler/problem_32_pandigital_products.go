package main

import (
	"fmt"
	"math"
)

/*
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/

func digits(num int) int {
	return int(math.Log10(float64(num))) + 1
}

func check(num int, digits []int) bool {

	for num != 0 {

		digit := num % 10
		if digits[digit] == 1 {
			return false
		}

		digits[digit] = 1
		num /= 10
	}

	return true
}

func isPandigital(a int, b int, c int) bool {

	digits := make([]int, 10)
	digits[0] = 1

	if !check(a, digits) || !check(b, digits) || !check(c, digits) {
		return false
	}

	return true
}

func prob32() int {

	var result int
	pandigitalMap := make(map[int]bool)

	for i := 1; i < 100; i++ {
		for j := i + 1; ; j++ {

			allDigits := digits(i) + digits(j) + digits(i*j)

			if allDigits > 9 {
				break
			}

			if allDigits == 9 && isPandigital(i, j, i*j) && !pandigitalMap[i*j] {
				result += i * j
				pandigitalMap[i*j] = true
			}
		}
	}

	return result
}

func main() {
	fmt.Println(prob32())
}
