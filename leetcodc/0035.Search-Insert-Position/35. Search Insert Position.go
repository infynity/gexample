package leetcode

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	// 0 5   1 2 3 4 5 6
	for low <= high {
		//mid := low + (high-low)>>1
		mid := low + (high-low)/2

		if nums[mid] >= target {
			high = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}
