package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type (
	Node struct {
		val  int
		next *Node
	}

	MyHashSet struct {
		list []*Node
	}
)

const (
	length = 1000
)

/** Initialize your data structure here. */
func Constructor() MyHashSet {

	list := make([]*Node, length)
	return MyHashSet{
		list: list,
	}
}

func (this *MyHashSet) Add(key int) {

	// Hash func: key % length
	bucket := key % length

	if !this.Contains(key) {
		node := Node{
			val:  key,
			next: this.list[bucket],
		}
		this.list[bucket] = &node
	}
}

func (this *MyHashSet) Remove(key int) {

	// Hash func: key % length
	bucket := key % length

	p := this.list[bucket]
	if p == nil {
		return
	}

	if p.val == key {
		this.list[bucket] = p.next
		return
	}

	for p.next != nil {
		if p.next.val == key {
			tmp := p.next
			p.next = p.next.next
			tmp.next = nil
			return
		}
		p = p.next
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {

	// Hash func: key % length
	bucket := key % length

	p := this.list[bucket]
	for p != nil {
		if p.val == key {
			return true
		}
		p = p.next
	}

	return false
}
