package main

import "fmt"

/*
Given an m x n matrix. If an element is 0, set its entire row and column to 0. Do it in-place.

Follow up:

A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])

	bitMapOfM := make([]int, m)
	bitMapOfN := make([]int, n)

	for i, line := range matrix {
		for j := range line {
			if matrix[i][j] == 0 {
				bitMapOfM[i] = 1
				bitMapOfN[j] = 1
			}
		}
	}

	for i := range matrix {

		if bitMapOfM[i] == 1 {
			for j := range matrix[i] {
				matrix[i][j] = 0
			}
			continue
		}

		for j := range matrix[i] {
			if bitMapOfN[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
