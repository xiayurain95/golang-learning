package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("{[}"))
}
func isValid(s string) bool {
	var tmp []byte
	m := map[byte]byte{'(': ')', '{': '}', '[': ']', ')': '(', ']': '[', '}': '{'}
	top := -1
	for i := 0; i < len(s); i++ {
		if top == -1 {
			top++
			tmp = append(tmp, s[i])
		} else {
			if tmp[top] == m[s[i]] {
				tmp = tmp[:len(tmp)-1]
				top--
			} else {
				top++
				tmp = append(tmp, s[i])
			}
		}
	}
	if top == -1 {
		return true
	} else {
		return false
	}
}
