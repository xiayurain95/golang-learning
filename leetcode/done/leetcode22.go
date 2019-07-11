package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}
func generateParenthesis(n int) []string {
	var outTmp [][]byte
	var solve func([]byte, int)
	var chkAns func([]byte) bool
	solve = func(b []byte, n int) {
		if n == 0 {
			if chkAns(b) {
				outTmp = append(outTmp, b)
			} else {
			}
		} else {
			tmpB := make([]byte, len(b))
			copy(tmpB, b)
			solve(append(tmpB, '('), n-1)
			solve(append(b, ')'), n-1)
		}
	}
	chkAns = func(ans []byte) bool {
		var tmp []byte
		cnt := -1
		for i := 0; i < len(ans); i++ {
			if cnt == -1 {
				cnt++
				tmp = append(tmp, ans[i])
			} else {
				if tmp[cnt] == '(' && ans[i] == ')' {
					cnt--
					tmp = tmp[:len(tmp)-1]
				} else {
					cnt++
					tmp = append(tmp, ans[i])
				}
			}
		}
		if cnt == -1 {
			return true
		} else {
			return false
		}
	}
	var tmp []byte
	solve(tmp, 2*n)
	var out []string
	for _, v := range outTmp {
		out = append(out, string(v))
	}
	return out
}
