package main

import "fmt"

func main() {
	setZeroes([][]int{{1, 1, 1}, {0, 1, 0}, {1, 1, 1}})
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m := len(matrix)
	n := len(matrix[0])
	row := make([]bool, m)
	col := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	fmt.Println(matrix)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(row)
	fmt.Println(col)

	for i := 0; i < m; i++ {
		if row[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < n; j++ {
		if col[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Println(matrix)
}
