package main

import (
	"fmt"
	"sort"
)

type pair struct {
	key   int
	value int
}

func topKFrequent(nums []int, k int) []int {
	var result []int
	intMap := make(map[int]int)
	for _, v := range nums {
		intMap[v]++
	}

	// fmt.Println(intMap)

	pairs := make([]pair, 0, len(intMap))
	for k, v := range intMap {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})

	// fmt.Println(pairs)

	for _, p := range pairs[:k] {
		result = append(result, p.key)
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}

	topFreq := topKFrequent(nums, 2)
	fmt.Println(topFreq)
}
