package main

/*
func main() {
	fmt.Println(myAtoi("20000000000000000000"))
}


func myAtoi(str string) int {
	target, flagNeg := func(str string) (string, bool) {
		type status struct {
			pre        bool
			neg        bool
			preZeroNum bool
			strCnt     int
			out        []byte
		}
		sta := status{true, false, true, 0, make([]byte, 0)}
		for sta.strCnt < len(str) {
			if sta.pre {
				if str[sta.strCnt] == ' ' {
					sta.strCnt++
				} else if str[sta.strCnt] == '+' {
					sta.strCnt++
					sta.pre = false
				} else if str[sta.strCnt] == '-' {
					sta.neg = true
					sta.pre = false
					sta.strCnt++
				} else {
					sta.pre = false
				}
			} else if !sta.pre && sta.preZeroNum {
				if str[sta.strCnt] == '0' {
					sta.strCnt++
				} else {
					sta.preZeroNum = false
				}
			} else if !sta.pre && !sta.preZeroNum {
				if str[sta.strCnt] >= '0' && str[sta.strCnt] <= '9' {
					sta.out = append(sta.out, str[sta.strCnt])
					sta.strCnt++
				} else {
					break
				}
			}
		}
		return string(sta.out), sta.neg
	}(str)
	if target == "" {
		return 0
	} else if flagNeg && (len(target) > len("2147483648") || (len(target) == len("2147483648") && target >= "2147483648")) {
		return -2147483648
	} else if !flagNeg && (len(target) > len("2147483647") || (len(target) == len("2147483647") && target >= "2147483647")) {
		return 2147483647
	} else {
		out, loops := 0, 1
		for i := len(target) - 1; i >= 0; i-- {
			out += int(target[i]-'0') * loops
			loops *= 10
		}
		if flagNeg {
			return -out
		} else {
			return out
		}
	}
}
*/
/*var target string
j, k, flagJ, flag, flagNeg := 0, 0, false, false, false
i := 0
for ; i < len(str); i++ {
	if str[i] == ' ' || str[i] == '+' || str[i] == '-' || str[i] == '0' {
		if (str[i] == '+' || str[i] == '-' || str[i] == ' ') && !flag {
			flag = !flag
			if str[i] == '-' {
				flagNeg = !flagNeg
			}
		} else if (str[i] == '+' || str[i] == '-' || str[i] == ' ') && flag {
			//k = i
			break
		} else if str[i] == ' ' && flagJ {
			break
		} else {
			continue
		}
	} else if '0' <= str[i] && str[i] <= '9' {
		if !flagJ {
			j = i
			flagJ = !flagJ
		} else {
			continue
		}

	} else {
		//k = i
		break
	}
}
k = i

if flagJ {
	target = str[j:k]
} else {
	target = ""
}
*/
