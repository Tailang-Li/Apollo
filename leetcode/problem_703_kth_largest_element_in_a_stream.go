package main

import "fmt"

/*
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
int add(int val) Returns the element representing the kth largest element in the stream.
*/

func main() {

}

type KthLargest struct {
	length int
	queue  []int
}

func Constructor(k int, nums []int) KthLargest {

	kthLargest := KthLargest{
		queue: make([]int, k),
	}

	for _, n := range nums {
		kthLargest.Add(n)
	}

	return kthLargest
}

func (this *KthLargest) Add(val int) int {

    if 
}
