package main

import "fmt"

/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/

const (
	maxInt = 2147483647
	minInt = -2147483648
)

func main() {
	fmt.Println(reverse(123456))
}

func reverse(x int) int {
	var result int
	for x != 0 {
		tmp := x % 10
		x = x / 10
		if result > maxInt/10 || (result == maxInt/10 && tmp > 7) {
			return 0
		}
		if result < minInt/10 || (result == minInt/10 && tmp < -8) {
			return 0
		}
		result = result*10 + tmp
	}
	return result
}
