package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(5))
}
func countAndSay(n int) string {
	out := "1"
	var solve func(int)
	solve = func(n int) {
		if n == 1 {
			return
		} else {
			var b []byte
			for cnt, i, j := 0, 0, 0; i < len(out); {
				if out[i] == out[j] {
					cnt++
					i++
					if i == len(out) {
						b = append(b, byte(cnt+int('0')))
						b = append(b, out[j])
					}
				} else {
					b = append(b, byte(cnt+int('0')))
					b = append(b, out[j])
					j = i
					cnt = 0
				}
			}
			out = string(b)
			solve(n - 1)
		}
	}
	solve(n)
	return out
}
