package main

import (
	"fmt"
)

/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

var arr [10000000]int64

/*
用数组存储中间结果即可，注意数组的长度应当设大一些，中间的某个值可能比一百万大
*/
func longestChan(n int64) int64 {

	if n == int64(1) {
		return 1
	}

	if n < 10000000 && arr[n] != 0 {
		return arr[n]
	}

	var tmp int64
	if n%2 == 0 {
		tmp = longestChan(n/2) + 1
	} else {
		tmp = longestChan(3*n+1) + 1
	}

	if n < 10000000 {
		arr[n] = tmp
	}

	return tmp
}

func main() {

	var (
		max   int64
		index int64
	)
	for i := int64(1); i <= int64(1000000); i++ {
		tmp := longestChan(i)
		if tmp > max {
			max = tmp
			index = i
		}
	}

	fmt.Println(index)
}
