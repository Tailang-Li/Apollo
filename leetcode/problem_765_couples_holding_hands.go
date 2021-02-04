package main

import "fmt"

func main() {
	fmt.Println(minSwapsCouples([]int{0, 2, 4, 6, 7, 1, 3, 5}))
}

func minSwapsCouples(row []int) int {

	indexMap := make([]int, len(row))
	for i, item := range row {
		indexMap[item] = i
	}

	var count int
	for i := 0; i < len(row); i += 2 {

		anotherOne := row[i] ^ 1
		if indexMap[anotherOne] == i+1 {
			continue
		}

		indexMap[row[i+1]] = indexMap[anotherOne]
		row[indexMap[anotherOne]] = row[i+1]
		count++
	}

	return count
}
