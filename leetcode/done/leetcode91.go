package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numDecodings("226"))
}

/*
func numDecodings(s string) int {
	m := map[string]int{}
	for i := 1; i <= 26; i++ {
		m[strconv.Itoa(i)] = 0
	}
	var solve func(string) int
	solve = func(str string) int {
		if len(str) == 1 {
			if _, ok := m[str[:1]]; ok {
				return 1
			} else {
				return 0
			}
		} else if len(str) == 0 {
			return 1
		} else {
			if str[0] == '0' {
				return 0
			} else {
				if _, ok := m[str[:2]]; ok {
					return solve(str[1:]) + solve(str[2:])
				} else {
					return solve(str[1:])
				}
			}
		}

	}
	return solve(s)
}

map里面使用string实在太浪费了,没办法,之后记得转换成int吧...
*/
func numDecodings(s string) int {
	m := map[string]int{}
	for i := 1; i <= 26; i++ {
		m[strconv.Itoa(i)] = 0
	}
	var solve func([]byte) int
	b := []byte(s)
	solve = func(str []byte) int {
		if len(str) == 1 {
			if str[0] == '0' {
				return 0
			} else {
				return 1
			}
		} else if len(str) == 0 {
			return 1
		} else {
			if str[0] == '0' {
				return 0
			} else {
				if _, ok := m[string(str[:2])]; ok {
					return solve(str[1:]) + solve(str[2:])
				} else {
					return solve(str[1:])
				}
			}
		}

	}
	return solve(b)
}
