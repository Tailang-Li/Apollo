package main

import "fmt"

func main() {
	fmt.Println(isToeplitzMatrix([][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}))
}

func isToeplitzMatrix(matrix [][]int) bool {

	rowLength := len(matrix)
	columeLength := len(matrix[0])

	// row
	for i := 0; i < rowLength; i++ {
		p, q, target := i, 0, matrix[i][0]
		for p < rowLength && q < columeLength {
			fmt.Println(matrix[p][q])
			if matrix[p][q] != target {
				return false
			}
			p++
			q++
		}
	}

	// colume
	for i := 1; i < len(matrix[0]); i++ {
		p, q, target := 0, i, matrix[0][i]
		for p < rowLength && q < columeLength {
			fmt.Println(matrix[p][q])
			if matrix[p][q] != target {
				return false
			}
			p++
			q++
		}
	}

	return true
}
