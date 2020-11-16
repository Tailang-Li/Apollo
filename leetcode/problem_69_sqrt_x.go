package leetcode

/*
Given a non-negative integer x, compute and return the square root of x.

Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.
*/

func mySqrt(x int) int {

	l, r := 0, x
	for l != r {
		mid := (l + r + 1) / 2
		if mid*mid <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}
