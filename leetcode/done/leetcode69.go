package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(1))
}
func mySqrt(x int) int {
	out := 0
	for out*out <= x {
		out++
	}
	out--
	return out
}
