package main

import (
	"fmt"
	"sort"
)

// func minDifference(nums []int) int {
// 	if len(nums) <= 4 {
// 		return 0
// 	}
// 	sort.Ints(nums)
// 	minDiff := nums[len(nums)-1] - nums[0]
// 	for i := 0; i <= 3; i++ {
// 		minDiff = min(minDiff, nums[len(nums)-(3-i)-1]-nums[i])
// 	}

// 	return minDiff
// }

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	sort.Ints(nums)
	ans := nums[len(nums)-1] - nums[0]
	for i := 0; i <= 3; i++ {
		ans = min(ans, nums[len(nums)-(3-i)-1]-nums[i])
	}
	return ans
}

func main() {
	// arr := []int{3, 199, 299}
	// arr := []int{5,3,2,4}
	arr := []int{1, 5, 0, 10, 14}
	// arr := []int{6, 6, 0, 1, 1, 4, 6}

	ans := minDifference(arr)
	fmt.Println(ans)
}
