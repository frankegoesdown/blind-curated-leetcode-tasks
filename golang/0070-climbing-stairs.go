package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(10))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	s1 := 1
	s2 := 2
	ret := 0
	for i := 3; i <= n; i++ {
		ret = s1 + s2
		s1, s2 = s2, ret
		fmt.Println(i)
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(ret)
		fmt.Println("====")
	}
	return ret
}
