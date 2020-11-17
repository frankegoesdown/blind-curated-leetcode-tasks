package main

import "fmt"

func main() {
	fmt.Println([]int{2, 7, 11, 15}, 9) // result [0, 1]
}

func twoSum(nums []int, target int) (result []int) {
	numbersMap := make(map[int]int)

	for idx, i := range nums {
		if j, ok := numbersMap[target-i]; ok {
			return []int{j, idx}
		}
		numbersMap[i] = idx
	}

	return
}
