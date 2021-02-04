package main

import "fmt"

/*
Given an integer array arr, return the length of a maximum size turbulent subarray of arr.

A subarray is turbulent if the comparison sign flips between each adjacent pair of elements in the subarray.

More formally, a subarray [arr[i], arr[i + 1], ..., arr[j]] of arr is said to be turbulent if and only if:

For i <= k < j:
    arr[k] > arr[k + 1] when k is odd, and
    arr[k] < arr[k + 1] when k is even.
Or, for i <= k < j:
    arr[k] > arr[k + 1] when k is even, and
    arr[k] < arr[k + 1] when k is odd.
*/

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
	fmt.Println(maxTurbulenceSize([]int{4, 8, 12, 16}))
}

func maxTurbulenceSize(arr []int) int {

	switch len(arr) {
	case 0, 1:
		return len(arr)
	}

	var (
		maxLength int
		i         int
		j         int
	)

LOOP:
	for i < len(arr)-1 {

		var tag bool
		if arr[i] > arr[i+1] || arr[i] < arr[i+1] {

			if arr[i] < arr[i+1] {
				tag = true
			}

			for j = i + 1; j < len(arr)-1; j++ {

				switch {
				case tag && arr[j] <= arr[j+1],
					!tag && arr[j] >= arr[j+1]:
					maxLength = max(maxLength, j-i+1)
					i = j
					continue LOOP
				}

				tag = changeTag(tag)
			}

			return max(maxLength, j-i+1)
		}

		i++
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func changeTag(tag bool) bool {
	if tag {
		return false
	}
	return true
}
