package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if i != 0 && digits[i]+1 == 10 {
			digits[i] = 0
		} else if i == 0 && digits[i]+1 == 10 {
			digits[i] = 0
			digits = append(digits, 0)
			copy(digits[1:], digits[:len(digits)-1])
			digits[0] = 1
			break
		} else {
			digits[i] += 1
			break
		}
	}
	return digits
}
