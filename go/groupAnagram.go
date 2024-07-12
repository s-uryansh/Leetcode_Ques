package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		anagramMap[count] = append(anagramMap[count], s)
	}
	// fmt.Println(anagramMap)
	groupedAnagrams := make([][]string, len(anagramMap))
	i := 0
	for _, v := range anagramMap {
		groupedAnagrams[i] = v
		i++
	}

	return groupedAnagrams
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	groupedAnagrams := groupAnagrams(strs)
	fmt.Println(groupedAnagrams)

	// twoDarry := [][]string{
	// 	{"tan", "nat"},
	// 	{"eat", "tea", "ate"},
	// }
	// fmt.Println(twoDarry)
}
