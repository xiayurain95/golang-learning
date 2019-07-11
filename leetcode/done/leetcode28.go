package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("hello", "ll"))
}

// vesion1
/*
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	} else {
		i, j := 0, -1
		for ; i < len(haystack); i++ {
			for k := i; k < len(haystack) && j+1 < len(needle) && haystack[k] == needle[j+1]; k++ {
				j++
			}
			if j+1 == len(needle) {
				break
			} else {
				j = -1
			}
		}
		if j != -1 {
			return i
		} else {
			return -1
		}
	}
}
*/
// for future maybe~~
func strStr(haystack string, needle string) int {

}
