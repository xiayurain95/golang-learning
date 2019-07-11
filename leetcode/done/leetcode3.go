package main

/*
import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	stringLen := len(s)
	ans := 0
	for i := 0; i < stringLen; i++ {
		tmp := make(map[byte]int)
		// flag := true
		cnt := 0
		for j := i; j < stringLen; j++ {
			_, ok := tmp[s[j]]
			if !ok {
				tmp[s[j]] = 1
				cnt++
			} else {
				break
			}
		}
		if ans < cnt {
			ans = cnt
		}
	}
	return ans
}
*/
