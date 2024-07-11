package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	flag := 0
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			flag += 1
		}

	}
	if flag != 0 {
		return true
	}
	return false
}

func main() {
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	isDupe := containsDuplicate(arr)
	fmt.Println(isDupe)
}
