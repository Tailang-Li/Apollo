package main

import "fmt"

/*
Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.
*/

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	count := 0
	direction := 0
	p, q := 0, -1
	step := 2 * n

	for {
		switch direction {
		case 0:
			for i := 0; i < step/2; i++ {
				q++
				result[p][q] = count + 1
				count++
			}
		case 1:
			for i := 0; i < step/2; i++ {
				p++
				result[p][q] = count + 1
				count++
			}
		case 2:
			for i := 0; i < step/2; i++ {
				q--
				result[p][q] = count + 1
				count++
			}
		case 3:
			for i := 0; i < step/2; i++ {
				p--
				result[p][q] = count + 1
				count++
			}
		}

		if count == n*n {
			return result
		}

		step--
		direction = (direction + 1) % 4
	}
}
