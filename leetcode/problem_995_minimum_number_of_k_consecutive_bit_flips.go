package main

import "fmt"

/*
In an array A containing only 0s and 1s, a K-bit flip consists of choosing a (contiguous) subarray of length K and simultaneously changing every 0 in the subarray to 1, and every 1 in the subarray to 0.

Return the minimum number of K-bit flips required so that there is no 0 in the array.  If it is not possible, return -1.
*/

/*
两个思路——差分数组和滑动窗口

前置思考：
对于最终解，翻动的顺序并不影响结果，所以可以认为在遇到 0 之后就立即翻，将子数组内的元素全部异或
即朴素贪心的方式，时间复杂度为 O(NK)

1. 差分数组
根据上面的朴素贪心的方式，其实并不需要每次遇到 0 之后都真正去翻一次。可以用差分数组来记录 A[i] 和 A[i-1] 之间翻动次数的差

例：
[0, 0, 0, 1, 0, 1, 1, 0]
翻第一次之后是：
[1, 1, 1, 1, 0, 1, 1, 0] => 对于子数组内而言，一起翻动的，那翻动次数的差肯定不变，变得一定是 A[i] 和 A[i+K] 这两个位置
那么我们在遍历到 A[0] 的时候，记录当前翻动的次数为 1，将 Diff[3] 设为 -1，表示 0 - 2 这个子数组翻过一次了
所以根据这个思路，首先在遍历到一个位置时，先用记录的当前翻动次数 + Diff 的值，就可以知道当前这个位置翻动了多少次
然后根据翻动的次数 + 原始数组中的值 % 2 就可以知道现在的实际值是多少
所以时间复杂度为 O(N)

2. 滑动窗口
在根据上面差分数组的思路，
*/

func main() {
	fmt.Println(simpleGreedy([]int{0, 0, 0, 1, 0, 1, 1, 0}, 3))
	fmt.Println(diffArray([]int{0, 0, 0, 1, 0, 1, 1, 0}, 3))
}

func simpleGreedy(A []int, K int) int {

	var count int
	for i := 0; i <= len(A)-K; i++ {
		if A[i] == 0 {
			count++
			for j := 1; j < K; j++ {
				A[i+j] ^= 1
			}
		}
	}

	for i := len(A) - K + 1; i < len(A); i++ {
		if A[i] == 0 {
			return -1
		}
	}

	return count
}

func diffArray(A []int, K int) int {

	var (
		count   int
		curFlip int
		length  int = len(A)
	)

	diff := make([]int, length+1)
	for i := range A {
		curFlip += diff[i]
		curNum := (curFlip + A[i]) % 2
		if curNum == 0 {
			if i+K > length {
				return -1
			}
			count++
			curFlip++
			diff[i+K]--
		}
	}

	return count
}
