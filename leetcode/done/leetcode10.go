package main

import "fmt"

func main() {
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("ab", ".*"), ",true")
	fmt.Println(isMatch("ab", ".*c"))
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("ab", ".*c"))
	fmt.Println(isMatch("aa", "a*"), ",true")
	fmt.Println(isMatch("aab", "c*a*b*"), ",true")
	fmt.Println(isMatch("aaa", "a*a"), ",true")
	fmt.Println(isMatch("aaa", "ab*a*c*a"), ",true")
	fmt.Println(isMatch("mississippi", "mis*is*ip*."), ",true")
	fmt.Println(isMatch("bbbba", ".*a*a"), ",true")
	fmt.Println(isMatch("a", "ab*"), ",true")
}

func isMatch(s string, p string) bool {
	if len(p) > 1 && p[1] == '*' {
		if len(s) > 0 && (p[0] == '.' || p[0] == s[0]) {
			return isMatch(s[1:], p[2:]) || isMatch(s[1:], p) || isMatch(s, p[2:])
		} else {
			return isMatch(s, p[2:])
		}
	} else {
		if len(s) > 0 && len(p) > 0 && (p[0] == '.' || p[0] == s[0]) {
			return isMatch(s[1:], p[1:])
		} else if len(s) == 0 && len(p) == 0 {
			return true
		} else {
			return false
		}
	}
}
