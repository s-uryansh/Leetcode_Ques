package main

import (
	"fmt"
)

func romanToInt(s string) int {
	var intValue int
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i, v := range s {
		intValue += roman[string(v)]
		if i != 0 {
			if roman[string(s[i-1])] < roman[string(v)] {
				intValue -= 2 * roman[string(s[i-1])]
			}
		}
	}
	return intValue
}
func main() {
	// s := "MCMXCIV"
	s := "LVIII"

	intVal := romanToInt(s)
	fmt.Println(intVal)
}
