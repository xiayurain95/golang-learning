package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{2, 5, 1, 7, 8, 43, -1, -49}))
}

func mergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	} else {
		mid := len(input) / 2
		bow := mergeSort(input[:mid])
		stern := mergeSort(input[mid:])
		return merge(bow, stern)
	}

}
func merge(bow, stern []int) []int {
	result := make([]int, 0, len(bow)+len(stern))
	i, j := 0, 0
	for i < len(bow) && j < len(stern) {
		if bow[i] < stern[j] {
			result = append(result, bow[i])
			i++
		} else {
			result = append(result, stern[j])
			j++
		}
	}
	result = append(result, bow[i:]...)
	result = append(result, stern[j:]...)
	return result
}
