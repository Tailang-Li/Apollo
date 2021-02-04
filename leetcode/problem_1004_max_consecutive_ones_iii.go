package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}

func longestOnes(A []int, K int) int {

	var (
		maxWindowSize int
		zeroCnt       int
		p             int
		q             int
	)

	for p < len(A) {

		if A[p] == 0 {
			zeroCnt++
		}

		for zeroCnt > K {

			if A[q] == 0 {
				zeroCnt--
			}

			q++
		}

		maxWindowSize = max(maxWindowSize, p-q+1)
		p++
	}

	return maxWindowSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
