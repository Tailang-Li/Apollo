package main

import "fmt"

/*
An array is monotonic if it is either monotone increasing or monotone decreasing.

An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].

Return true if and only if the given array A is monotonic.
*/

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
	fmt.Println(isMonotonic([]int{1, 3, 2}))
}

func isMonotonic(A []int) bool {

	if len(A) <= 2 {
		return true
	}

	greater := -1
	for i := 1; i < len(A); i++ {

		if A[i] == A[i-1] {
			continue
		}

		if greater == -1 {
			greater = 0
			if A[i] > A[i-1] {
				greater = 1
			}
			continue
		}

		if greater == 1 && A[i] < A[i-1] ||
			greater == 0 && A[i] > A[i-1] {
			return false
		}
	}

	return true
}
