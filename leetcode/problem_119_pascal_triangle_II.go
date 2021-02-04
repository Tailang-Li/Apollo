package main

import "fmt"

/*
Given an integer rowIndex, return the rowIndexth row of the Pascal's triangle.

Notice that the row index starts from 0.
*/

func main() {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {

	result := make([]int, rowIndex+1)

	result[0] = 1
	for i := 1; i < len(result); i++ {

		result[i] = 1
		for j := i - 1; j >= 1; j-- {
			result[j] = result[j-1] + result[j]
		}
	}

	return result
}
