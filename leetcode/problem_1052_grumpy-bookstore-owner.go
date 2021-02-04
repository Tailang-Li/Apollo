package main

import "fmt"

/*
Today, the bookstore owner has a store open for customers.length minutes.  Every minute, some number of customers (customers[i]) enter the store, and all those customers leave after the end of that minute.

On some minutes, the bookstore owner is grumpy.  If the bookstore owner is grumpy on the i-th minute, grumpy[i] = 1, otherwise grumpy[i] = 0.  When the bookstore owner is grumpy, the customers of that minute are not satisfied, otherwise they are satisfied.

The bookstore owner knows a secret technique to keep themselves not grumpy for X minutes straight, but can only use it once.

Return the maximum number of customers that can be satisfied throughout the day.
*/

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
	fmt.Println(maxSatisfied([]int{3}, []int{1}, 1))
}

func maxSatisfied(customers []int, grumpy []int, X int) int {

	var originalCnt int
	for i := range customers {
		if grumpy[i] == 0 {
			originalCnt += customers[i]
		}
	}

	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			originalCnt += customers[i]
		}
	}

	var (
		maxCnt int = originalCnt
	)
	for i := X; i < len(customers); i++ {

		if grumpy[i] == 1 {
			originalCnt += customers[i]
		}
		if grumpy[i-X] == 1 {
			originalCnt -= customers[i-X]
		}

		maxCnt = max(maxCnt, originalCnt)
	}

	return maxCnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
