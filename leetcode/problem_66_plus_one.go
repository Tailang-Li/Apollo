package main

import "fmt"

/*
Given a non-empty array of decimal digits representing a non-negative integer, increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.
*/

func main() {
	fmt.Println(plusOne([]int{9, 9}))
}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}

	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}
