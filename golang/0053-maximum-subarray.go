package main

func main() {

}

func maxSubArray(nums []int) int {
	maxSum, currSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
