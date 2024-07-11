package main

import (
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {

	arrS := arrCreate(s)
	arrT := arrCreate(t)
	sort.Strings(arrS)
	sort.Strings(arrT)

	if len(s) != len(t) {
		return false
	}
	flag := 0

	for i := 0; i < len(arrS); i++ {
		if arrS[i] == arrT[i] {
			continue
		} else {
			flag++
		}
	}
	if flag != 0 {
		return false
	}

	return true
}

func arrCreate(varString string) []string {

	arr := make([]string, len(varString))
	for i, v := range varString {
		arr[i] = string(v)
	}
	return arr
}

func main() {
	s := "anagram"
	t := "nagaram"

	isValidAnagram := isAnagram(s, t)
	fmt.Println(isValidAnagram)

}
