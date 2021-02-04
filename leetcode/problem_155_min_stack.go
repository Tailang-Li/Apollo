package main

import "fmt"

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
*/

func main() {
	minStack := Constructor()
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	fmt.Println(minStack.Top())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	minStack.Push(2147483647)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Push(-2147483648)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
}

type MinStack struct {
	stack []int64
	min   int64
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int64, 0),
	}
}

func (this *MinStack) Push(x int) {

	if len(this.stack) == 0 {
		this.min = int64(x)
	}

	this.stack = append(this.stack, int64(x)-int64(this.min))

	if int64(x) < this.min {
		this.min = int64(x)
	}
}

func (this *MinStack) Pop() {

	if len(this.stack) == 0 {
		return
	}

	if this.stack[len(this.stack)-1] < 0 {
		this.min = this.min - this.stack[len(this.stack)-1]
	}

	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {

	if len(this.stack) == 0 {
		return 0
	}

	if this.stack[len(this.stack)-1] < 0 {
		return int(this.min)
	}

	return int(this.min + this.stack[len(this.stack)-1])
}

func (this *MinStack) GetMin() int {
	return int(this.min)
}
