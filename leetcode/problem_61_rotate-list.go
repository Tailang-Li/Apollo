package main

import "fmt"

/*
Given the head of a linked list, rotate the list to the right by k places.
*/

func main() {
	fmt.Println("vim-go")
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil {
		return nil
	}

	p := head
	var cnt int
	for p.Next != nil {
		p = p.Next
		cnt++
	}
	p.Next = head
	cnt++

	k = k % cnt
	for i := 0; i < cnt-k; i++ {
		p = p.Next
	}

	result := p.Next
	p.Next = nil
	return result
}
