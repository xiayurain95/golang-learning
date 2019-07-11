package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(isPalindrome("0P"))
}
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	lowBytes := bytes.ToLower([]byte(s))
	for i, j := 0, len(lowBytes)-1; i < j; i, j = i+1, j-1 {
		for i < len(lowBytes) && (lowBytes[i] < 'a' || lowBytes[i] > 'z') && (lowBytes[i] < '0' || lowBytes[i] > '9') {
			i++
		}
		for j > -1 && (lowBytes[j] < 'a' || lowBytes[j] > 'z') && (lowBytes[j] < '0' || lowBytes[j] > '9') {
			j--
		}

		if i < len(lowBytes) && j > -1 {
			if lowBytes[i] != lowBytes[j] {
				return false
			}
		} else {
			if i > len(lowBytes)-1 && j < 0 {
				return true
			} else {
				return false
			}
		}
	}
	return true
}
