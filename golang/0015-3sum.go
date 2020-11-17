package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	ans, n := [][]int{}, len(nums)
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		low, high, sum := i+1, n-1, 0-nums[i]
		fmt.Println(low)
		fmt.Println(high)
		fmt.Println(sum)
		fmt.Println("==========")
		for low < high {
			if nums[low]+nums[high] == sum {
				ans = append(ans, []int{nums[i], nums[low], nums[high]})
				for low < high && nums[low] == nums[low+1] {
					low++
				}
				for low < high && nums[high] == nums[high-1] {
					high--
				}
				low++
				high--
			} else if nums[low]+nums[high] > sum {
				high--
			} else {
				low++
			}
		}
	}
	return ans
}
