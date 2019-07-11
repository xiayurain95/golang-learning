package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{}))
}
func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int)
	var out [][]string
	cnt := 0
	for _, str := range strs {
		tmpStr := []byte(str)
		sort.Sort(byteSlice(tmpStr))
		if num, ok := m[string(tmpStr)]; ok {
			out[num] = append(out[num], str)
		} else {
			m[string(tmpStr)] = cnt
			out = append(out, []string{str})
			cnt++
		}
	}
	return out
}

type byteSlice []byte

func (b byteSlice) Len() int           { return len(b) }
func (b byteSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byteSlice) Less(i, j int) bool { return b[i] < b[j] }
