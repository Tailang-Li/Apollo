package main

import "fmt"

/*
The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
*/

func isPalindromic(originNum int, base int) bool {

	r, num := 0, originNum
	for num != 0 {
		r = r*base + num%base
		num /= base
	}

	return r == originNum
}

func prob36() int {

	var result int
	for i := 1; i < 1000000; i++ {
		if isPalindromic(i, 2) && isPalindromic(i, 10) {
			result += i
		}
	}

	return result
}

func main() {
	fmt.Println(prob36())
}
