package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	if bottom == 0 {
		return matrix[0]
	}
	if right == 0 {
		for i := 0; i <= bottom; i++ {
			res = append(res, matrix[i][0])
		}
		return res
	}
	for left <= right || top <= bottom {
		// 上边
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}
		// 右边
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		// 下边
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}
		// 左边
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}
