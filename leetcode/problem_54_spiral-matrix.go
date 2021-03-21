package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}

func spiralOrder(matrix [][]int) []int {

	m := len(matrix)
	n := len(matrix[0])
	path := make([][]int, m)
	result := make([]int, 0, m*n)
	for i := range path {
		path[i] = make([]int, n)
	}
	direction := 0
	p, q := 0, 0

	for {
		switch direction {
		case 0:
			for q < n && path[p][q] == 0 {
				result = append(result, matrix[p][q])
				path[p][q] = 1
				q++
			}
			p++
			q--
		case 1:
			for p < m && path[p][q] == 0 {
				result = append(result, matrix[p][q])
				path[p][q] = 1
				p++
			}
			p--
			q--
		case 2:
			for q >= 0 && path[p][q] == 0 {
				result = append(result, matrix[p][q])
				path[p][q] = 1
				q--
			}
			p--
			q++
		case 3:
			for p >= 0 && path[p][q] == 0 {
				result = append(result, matrix[p][q])
				path[p][q] = 1
				p--
			}
			p++
			q++
		}

		direction = (direction + 1) % 4

		if len(result) == m*n {
			return result
		}
	}
}
