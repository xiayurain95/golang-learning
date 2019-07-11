package main

/*
import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("abcd"))
}
func longestPalindrome(s string) string {
	m := make(map[int]string)
	strLen := len(s)
	if strLen > 2 {
		for i := 1; i < strLen-1; i++ {
			solve := func(j, k int) {
				// j, k := i-1, i+1
				walking := true
				for j >= 0 && k < strLen && walking {
					walking = false
					if s[j] == s[k] {
						j -= 1
						k += 1
						walking = true
					} //else if s[k] == s[i] && firstRun {
					// k += 1
					// walking, firstRun = true, false
					// } else if s[j] == s[i] && firstRun {
					// j -= 1
					// walking, firstRun = true, false
					// }
				}
				m[len(s[j+1:k])] = s[j+1 : k]
			}
			if s[i-1] == s[i] {
				solve(i-2, i+1)
			}
			if s[i] == s[i+1] {
				solve(i-1, i+2)
			}
			if s[i-1] == s[i+1] {
				solve(i-2, i+2)
			}
		}

		tmp := 0
		for k := range m {
			if k > tmp {
				tmp = k
			}
		}
		if tmp == 0 {
			return string(s[0])
		} else {
			return m[tmp]
		}
	} else if strLen == 1 {
		return s
	} else if strLen == 2 {
		if s[0] != s[1] {
			return string(s[0])
		} else {
			return s
		}
	} else {
		return s
	}
}
*/
