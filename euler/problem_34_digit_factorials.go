package main

import "fmt"

/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: As 1! = 1 and 2! = 2 are not sums they are not included.
*/

func isSame(originNum int, factorials []int) bool {

	num, tmp := originNum, 0
	for num != 0 {
		tmp += factorials[num%10]
		num /= 10
	}

	return tmp == originNum
}

func prob34() int {

	factorials := make([]int, 10)
	factorials[0] = 1
	for i := 1; i < len(factorials); i++ {
		factorials[i] = factorials[i-1] * i
	}

	var result int
	for i := 3; i < 10000000; i++ {
		if isSame(i, factorials) {
			result += i
		}
	}

	return result
}

func main() {
	fmt.Println(prob34())
}
