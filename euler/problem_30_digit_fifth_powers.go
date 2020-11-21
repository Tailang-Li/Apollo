package main

import "fmt"

/*
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + 4^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

func isSame(originNum int, powers []int) bool {

	num, temp := originNum, 0
	for num != 0 {
		temp += powers[num%10]
		num /= 10
	}

	return temp == originNum
}

func prob30() int {

	powers := make([]int, 10)
	for i := range powers {

		power := 1
		for j := 0; j < 5; j++ {
			power *= i
		}

		powers[i] = power
	}

	var result int
	for i := 2; i <= 999999; i++ {
		if isSame(i, powers) {
			result += i
		}
	}

	return result
}

func main() {
	fmt.Println(prob30())
}
