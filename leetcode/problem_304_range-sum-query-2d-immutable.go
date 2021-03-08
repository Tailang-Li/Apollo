package main

import "fmt"

/*
Given a 2D matrix matrix, find the sum of the elements inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).

The above rectangle (with the red border) is defined by (row1, col1) = (2, 1) and (row2, col2) = (4, 3), which contains sum = 8.
*/

func main() {
	obj := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	fmt.Println(obj.SumRegion(2, 1, 4, 3))
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
	fmt.Println(obj.SumRegion(1, 2, 2, 4))

	Constructor([][]int{})
}

type NumMatrix struct {
	cntMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {

	numMatrix := NumMatrix{
		cntMatrix: make([][]int, len(matrix)+1),
	}

	if len(matrix) == 0 {
		return numMatrix
	}

	numMatrix.cntMatrix[0] = make([]int, len(matrix[0])+1)
	for i, line := range matrix {
		numMatrix.cntMatrix[i+1] = make([]int, len(matrix[0])+1)
		var lineCnt int
		for j, num := range line {
			numMatrix.cntMatrix[i+1][j+1] = numMatrix.cntMatrix[i][j+1] + lineCnt + num
			lineCnt += num
		}
	}

	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.cntMatrix[row2+1][col2+1] - this.cntMatrix[row1][col2+1] - this.cntMatrix[row2+1][col1] + this.cntMatrix[row1][col1]
}
