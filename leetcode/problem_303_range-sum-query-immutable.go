package main

import "fmt"

/*
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int i, int j) Return the sum of the elements of the nums array in the range [i, j] inclusive (i.e., sum(nums[i], nums[i + 1], ... , nums[j]))
*/

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}

type NumArray struct {
	counts []int
}

func Constructor(nums []int) NumArray {

	numArray := NumArray{
		counts: make([]int, len(nums)+1),
	}

	for i := range nums {
		numArray.counts[i+1] = numArray.counts[i] + nums[i]
	}

	return numArray
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.counts[j+1] - this.counts[i]
}
