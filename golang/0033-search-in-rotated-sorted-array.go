package main

func main() {

}

func search(nums []int, target int) int {
	low, mid, high := 0, 0, len(nums)-1
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[low] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
