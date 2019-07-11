package main

import (
	"fmt"
)

func main() {
	fmt.Println(partition("abb"))
}
func partition(s string) [][]string {
	var out [][]string
	if len(s) == 0 {
		return out
	}
	var isPalindrome func([]byte) bool
	isPalindrome = func(s []byte) bool {
		if len(s) == 0 {
			return true
		}
		for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	var solve func([]byte, []string)
	solve = func(b []byte, ans []string) {
		if len(b) == 0 {
			out = append(out, ans)
		}
		for i := 1; i <= len(b); i++ {
			if isPalindrome(b[:i]) {
				tmp := make([]string, len(ans))
				copy(tmp, ans)
				tmp = append(tmp, string(b[:i]))
				solve(b[i:], tmp)
			}
		}
	}
	solve([]byte(s), []string{})
	return out
}
