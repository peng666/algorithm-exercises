package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i := 0
	j := len(matrix[0]) - 1
	for i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i += 1
		} else {
			j -= 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	res := searchMatrix(matrix, 20)
	fmt.Println(res)
	res = searchMatrix(matrix, 6)
	fmt.Println(res)
}
