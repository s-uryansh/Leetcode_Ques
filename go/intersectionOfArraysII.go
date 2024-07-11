package main

import (
	"fmt"
	"sort"
)

func concatenate(arr1 []int, arr2 []int) []int {
	return append(arr1, arr2...)
}
func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func intersect(arr1 []int, arr2 []int) []int {
	var addedArr []int
	var intersectedArr []int

	sort.Ints(arr1)
	sort.Ints(arr2)
	addedArr = concatenate(arr1, arr2)
	sort.Ints(addedArr)
	// fmt.Println(addedArr)

	for i := 0; i < len(addedArr)-1; i++ {
		if addedArr[i] == addedArr[i+1] {
			if !contains(intersectedArr, addedArr[i]) {
				intersectedArr = append(intersectedArr, addedArr[i])
			}
		}
	}

	return intersectedArr
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{9, 2, 3, 89, 23}

	intersectedArr := intersect(arr1, arr2)
	fmt.Println(intersectedArr)
}
