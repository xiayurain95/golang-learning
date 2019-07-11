package main

/*
func main() {
	fmt.Println(convert("ABC", 2))
}
func convert(s string, numRows int) string {
	if numRows > 1 {
		strLen := len(s)
		var tmpOut [][]byte
		cnt := 0
		isZigZag, zigZagCnt := false, 1
		for cnt < strLen {
			tmp := make([]byte, numRows)
			for i := numRows - 1; i >= 0; i-- {
				if cnt < strLen {
					if !isZigZag || numRows == 2 {
						tmp[i] = s[cnt]
						cnt++
					} else {
						if zigZagCnt < numRows-1 {
							if zigZagCnt == i {
								tmp[i] = s[cnt]
								cnt++
							} else {
								tmp[i] = ' '
							}
						}
					}
				} else {
					tmp[i] = ' '
				}

			}
			if !isZigZag {
				isZigZag = !isZigZag
			} else {
				if zigZagCnt < numRows-2 {
					zigZagCnt++
				} else {
					zigZagCnt = 1
					isZigZag = !isZigZag
				}
			}
			tmpOut = append(tmpOut, tmp)
		}
		out, outCnt := make([]byte, len(s)), 0
		for i := numRows - 1; i >= 0; i-- {
			for j := 0; j < len(tmpOut); j++ {
				if tmpOut[j][i] != ' ' {
					out[outCnt] = tmpOut[j][i]
					outCnt++
				}
			}
		}
		return string(out)
	} else {
		return s
	}
}
*/
