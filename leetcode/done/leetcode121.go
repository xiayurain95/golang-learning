package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{2, 4, 1}))
}
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max, min := 0, 0
	m := make(map[int]int)
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[min] {
			min, max = i, i
		}
		if prices[i] > prices[max] {
			max = i
			m[min] = prices[max] - prices[min]
		}
	}
	var out int
	for _, v := range m {
		if v > out {
			out = v
		}
	}
	return out
}
