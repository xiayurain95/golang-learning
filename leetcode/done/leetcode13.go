package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
func romanToInt(s string) int {
	numMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	abnormalNumMap := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	var out int
	for i := 0; i < len(s); {
		if i+1 <= len(s)-1 {
			tmp, ok := abnormalNumMap[s[i:i+2]]
			if ok {
				out += tmp
				i += 2
			} else {
				out += numMap[s[i]]
				i += 1
			}
		} else {
			out += numMap[s[i]]
			i += 1
		}

	}
	return out
}
