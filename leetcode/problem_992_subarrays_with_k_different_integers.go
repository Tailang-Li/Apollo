package main

import "fmt"

/*
Given an array A of positive integers, call a (contiguous, not necessarily distinct) subarray of A good if the number of different integers in that subarray is exactly K.

(For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.)

Return the number of good subarrays of A.
*/

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3))
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
	fmt.Println(atMostKDistinct([]int{1, 2, 1, 3, 4}, 3))
	fmt.Println(atMostKDistinct([]int{1, 2, 1, 3, 4}, 2))
	fmt.Println(atMostKDistinct([]int{1, 2, 1, 2, 3}, 2))
	fmt.Println(atMostKDistinct([]int{1, 2, 1, 2, 3}, 1))
}

func subarraysWithKDistinct(A []int, K int) int {
	return atMostKDistinct(A, K) - atMostKDistinct(A, K-1)
}

// 集合问题，恰好K个不同整数可以理解为，最多K个整数的集合 - 最多K-1个整数的集合
func atMostKDistinct(A []int, K int) int {

	var (
		left, right, result, count int
		freq                       []int = make([]int, len(A)+1)
	)

	for right < len(A) {

		if freq[A[right]] == 0 {
			count++
		}
		freq[A[right]]++
		right++

		for count > K {
			freq[A[left]]--
			if freq[A[left]] == 0 {
				count--
			}
			left++
		}

		result += right - left
	}

	return result
}
