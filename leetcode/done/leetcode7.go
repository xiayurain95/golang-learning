package main

/*
func main() {
	fmt.Println(reverse(-120))
}
func reverse(x int) int {
	// overflow := false
	flag := false
	var out string
	if x < 0 {
		out = strconv.Itoa(-x)
		flag = true
	} else {
		out = strconv.Itoa(x)
	}
	data := []rune(out)
	l := len(data)
	var tmp rune
	for i := 0; i < l/2; i++ {
		tmp = data[i]
		data[i] = '0'
		data[i] = data[l-1-i]
		data[l-1-i] = tmp
	}
	res := string(data)
	if !flag && len(res) >= len(strconv.Itoa(math.MaxInt32)) && res > strconv.Itoa(math.MaxInt32) {
		return 0
	} else if flag && len(res) >= len(strconv.Itoa(math.MinInt32)[1:]) && res > strconv.Itoa(math.MinInt32)[1:] {
		return 0
	} else {
		output, _ := strconv.Atoi(res)
		if flag {
			return -output
		} else {
			return output

		}
	}
}
*/
