package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"a", "a"}))
}
func longestCommonPrefix(strs []string) string {
	var out []byte
	if len(strs) > 0 {
		for i := 0; i < len(strs[0]); i++ {
			flag := false
			for j := 1; j < len(strs); j++ {
				if i < len(strs[j]) && strs[0][i] == strs[j][i] {
					continue
				} else {
					flag = !flag
					break
				}
			}
			if !flag {
				out = append(out, strs[0][i])
			} else {
				break
			}
		}
	}
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else {
		return string(out)
	}
}
