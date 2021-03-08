package main

import "fmt"

/*
Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
int pop() Removes the element from the front of the queue and returns it.
int peek() Returns the element at the front of the queue.
boolean empty() Returns true if the queue is empty, false otherwise.
Notes:

You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
*/

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	fmt.Println(obj.Pop())
	fmt.Println(obj.inStack, obj.outStack)
	obj.Push(5)
	fmt.Println(obj.inStack, obj.outStack)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
}

type MyQueue struct {
	inStack  []int
	outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

	var result int
	if this.Empty() {
		return result
	}

	if len(this.outStack) > 0 {
		result = this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[:len(this.outStack)-1]
		return result
	}

	for j := len(this.inStack) - 1; j > 0; j-- {
		this.outStack = append(this.outStack, this.inStack[j])
	}

	result = this.inStack[0]
	this.inStack = []int{}

	return result
}

/** Get the front element. */
func (this *MyQueue) Peek() int {

	var result int
	if this.Empty() {
		return result
	}

	if len(this.outStack) > 0 {
		return this.outStack[len(this.outStack)-1]
	}

	return this.inStack[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}
