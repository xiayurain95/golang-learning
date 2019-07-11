package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(10))
}
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	out := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			out = false
		}
	}
	return out
}
