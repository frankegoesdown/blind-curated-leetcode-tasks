package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for idx, _ := range dp {
		dp[idx] = make([]int, n)
		dp[idx][n-1] = 1
	}
	fmt.Println(dp)
	for idx, _ := range dp[0] {
		dp[m-1][idx] = 1
	}
	fmt.Println(dp)
	fmt.Println("-----")
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
			fmt.Println(dp)
		}
	}
	return dp[0][0]
}

/*
[[0 0 0 0 0 0 1]     [0 0 0 0 0 2 1] [1 1 1 1 1 1 1]]
[[0 0 0 0 0 0 1]     [0 0 0 0 3 2 1] [1 1 1 1 1 1 1]]
[[0 0 0 0 0 0 1]     [0 0 0 4 3 2 1] [1 1 1 1 1 1 1]]
[[0 0 0 0 0 0 1]     [0 0 5 4 3 2 1] [1 1 1 1 1 1 1]]
[[0 0 0 0 0 0 1]     [0 6 5 4 3 2 1] [1 1 1 1 1 1 1]]
[[0 0 0 0 0 0 1]     [7 6 5 4 3 2 1] [1 1 1 1 1 1 1]]
[[0 0 0 0 0 3 1]     [7 6 5 4 3 2 1] [1 1 1 1 1 1 1]]
[[0 0 0 0 6 3 1]     [7 6 5 4 3 2 1] [1 1 1 1 1 1 1]]
[[0 0 0 10 6 3 1]    [7 6 5 4 3 2 1] [1 1 1 1 1 1 1]]
[[0 0 15 10 6 3 1]   [7 6 5 4 3 2 1] [1 1 1 1 1 1 1]]
[[0 21 15 10 6 3 1]  [7 6 5 4 3 2 1] [1 1 1 1 1 1 1]]
[[28 21 15 10 6 3 1] [7 6 5 4 3 2 1] [1 1 1 1 1 1 1]]
*/
