package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("234"))
}
func letterCombinations(digits string) []string {
	phoneNum := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	var out []string
	var tmpString [][]byte
	var solve func(b [][]byte, digits string) [][]byte
	solve = func(b [][]byte, digits string) [][]byte {
		if len(digits) == 0 {
			return b
		} else {
			lenB := len(b)
			if lenB != 0 {
				for i := 0; i < lenB; i++ {
					b[i] = append(b[i], phoneNum[digits[0]][0])
					for j := 1; j < len(phoneNum[digits[0]]); j++ {
						tmp := make([]byte, len(b[i])-1)
						copy(tmp, b[i][:len(b[i])-1])
						tmp = append(tmp, phoneNum[digits[0]][j])
						b = append(b, tmp)
					}
				}
			} else {
				for i := 0; i < len(phoneNum[digits[0]]); i++ {
					b = append(b, []byte{phoneNum[digits[0]][i]})
				}
			}
			return solve(b, digits[1:len(digits)])
		}
	}
	tmpString = solve(tmpString, digits)
	for _, v := range tmpString {
		out = append(out, string(v))
	}
	return out
}
