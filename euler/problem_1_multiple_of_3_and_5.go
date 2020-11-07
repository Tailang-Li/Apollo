package main

import "fmt"

/*
Problem 1 Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

/*
找到所有 3 的倍数和 5 的倍数之和，在减去 15 的倍数之和即可
*/
func Prob1() int {
	// 等差数列求解
	multipleOf3 := (3 + 999) * 333 / 2
	multipleOf5 := (5 + 995) * 199 / 2
	multipleOf15 := (15 + 990) * 66 / 2
	return multipleOf3 + multipleOf5 - multipleOf15
}

func main() {
	fmt.Println(Prob1())
}
