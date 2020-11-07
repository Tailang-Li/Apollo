package main

import "fmt"

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

/*
首先将当前数反转，在判断是否和原数相同即可判断是否是回文数
*/
func isPalindromic(num int) bool {

	tmp, origin := 0, num

	for num != 0 {
		tmp = tmp*10 + num%10
		num /= 10
	}

	return origin == tmp
}

func prob4() int {

	var max int

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			num := i * j
			if isPalindromic(num) && num > max {
				max = num
			}
		}
	}

	return max
}

func main() {
	fmt.Println(prob4())
}
