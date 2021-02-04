package main

import (
	"fmt"
	"strconv"
)

/*
Given two binary strings a and b, return their sum as a binary string.
*/

func main() {
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {

	maxLenght := max(len(a), len(b))
	result := ""
	carry := 0

	for i := 0; i < maxLenght; i++ {

		if i < len(a) {
			carry += int(a[len(a)-i-1] - '0')
		}

		if i < len(b) {
			carry += int(b[len(b)-i-1] - '0')
		}

		result = strconv.Itoa(carry%2) + result
		carry /= 2
	}

	if carry > 0 {
		return "1" + result
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
