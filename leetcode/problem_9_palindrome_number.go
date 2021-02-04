package main

import (
	"fmt"
)

/*
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.
*/

func main() {
	fmt.Println(isPalindrome(12321))
}

func isPalindrome(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var palindromeNum int
	for palindromeNum < x {
		palindromeNum = palindromeNum*10 + x%10
		x /= 10
	}

	if palindromeNum == x || palindromeNum == x*10+palindromeNum%10 {
		return true
	}

	return false
}
