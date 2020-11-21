package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func canJump(J []int) bool {
	maxJump := 0
	for i := 0; i < len(J) && maxJump >= i; i++ {
		maxJump = max(maxJump, i+J[i])
		if maxJump >= len(J)-1 {
			return true
		}
	}

	return false
}
