package main

import "fmt"

/*
Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.
*/

func main() {
	testNums := []int{1, 1, 2, 2, 3, 3}
	fmt.Println(removeElement(testNums, 2))
	fmt.Println(testNums)
}

func removeElement(nums []int, val int) int {

	var p int
	for _, v := range nums {
		if v == val {
			continue
		}
		nums[p] = v
		p++
	}

	return p
}
