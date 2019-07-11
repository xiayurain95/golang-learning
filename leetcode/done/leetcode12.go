package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(intToRoman(58))
}
func intToRoman(num int) string {
	numMap := []map[byte]string{
		map[byte]string{'0': "", '1': "I", '2': "II", '3': "III", '4': "IV", '5': "V", '6': "VI", '7': "VII", '8': "VIII", '9': "IX"},
		map[byte]string{'0': "", '1': "X", '2': "XX", '3': "XXX", '4': "XL", '5': "L", '6': "LX", '7': "LXX", '8': "LXXX", '9': "XC"},
		map[byte]string{'0': "", '1': "C", '2': "CC", '3': "CCC", '4': "CD", '5': "D", '6': "DC", '7': "DCC", '8': "DCCC", '9': "CM"},
		map[byte]string{'0': "", '1': "M", '2': "MM", '3': "MMM"}}
	strNum := strconv.Itoa(num)
	var strOut string
	for i := len(strNum) - 1; i >= 0; i-- {
		strOut += numMap[i][strNum[len(strNum)-1-i]]
	}
	return strOut
}
