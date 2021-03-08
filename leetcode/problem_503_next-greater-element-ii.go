package main

import "fmt"

/*
Given a circular array (the next element of the last element is the first element of the array),
print the Next Greater Number for every element. The Next Greater Number of a number x is the first
greater number to its traversing-order next in the array, which means you could search circularly to
find its next greater number. If it doesn't exist, output -1 for this number.
*/

func main() {
	fmt.Println(nextGreaterElements([]int{2, 2, 2, 2}))
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}

func nextGreaterElements(nums []int) []int {

	length := len(nums)
	stack := make([]int, 0)
	ans := make([]int, length)
	for i := range ans {
		ans[i] = -1
	}

	for i := 0; i < 2*len(nums)-1; i++ {

		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%length] {
			ans[stack[len(stack)-1]] = nums[i%length]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i%length)
	}

	return ans
}
