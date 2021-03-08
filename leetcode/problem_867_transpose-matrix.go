package main

import "fmt"

/*
Given a 2D integer array matrix, return the transpose of matrix.

The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.
*/

func main() {
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

func transpose(matrix [][]int) [][]int {

	m, n := len(matrix), len(matrix[0])
	result := make([][]int, n)

	for i := range result {
		result[i] = make([]int, m)
		for j := range result[i] {
			result[i][j] = matrix[j][i]
		}
	}

	return result
}
