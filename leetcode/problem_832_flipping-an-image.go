package main

import "fmt"

/*
Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].
*/

func main() {
	fmt.Println(flipAndInvertImage([][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}))
	fmt.Println(flipAndInvertImage([][]int{
		{1, 1, 0, 0},
		{1, 0, 0, 1},
		{0, 1, 1, 1},
		{1, 0, 1, 0},
	}))
}

func flipAndInvertImage(A [][]int) [][]int {

	var temp int
	for i := range A {

		lineLength := len(A[i])
		for j := 0; j < lineLength/2; j++ {

			if A[i][j] != A[i][lineLength-1-j] {
				continue
			}

			A[i][j] ^= 1
			A[i][lineLength-1-j] ^= 1
		}

		if lineLength%2 != 0 {
			A[i][lineLength/2] ^= 1
		}
	}

	return A
}
