package main

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
	var consecutiveCheck int

	for _, value := range arr {

		if value%2 == 0 {
			consecutiveCheck = 0
		} else {
			consecutiveCheck++
		}

		if consecutiveCheck >= 3 {
			return true
		}
	}

	return false

}

func main() {
	arr := []int{1, 3, 2}

	isConsecutiveOdd := threeConsecutiveOdds(arr)
	fmt.Println(isConsecutiveOdd)
	return
}
