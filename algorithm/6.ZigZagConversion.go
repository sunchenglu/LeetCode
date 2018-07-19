package algorithm

import "fmt"

func ZigZagConvert(s string, numRows int) string {

	length := len(s)

	if numRows == 1 || length == 1 {
		return s
	}

	zArr := make([][]byte, numRows)
	flag := true

	var result []byte

	for i:=0; i<length; i++ {

		rest := i%(numRows-1)

		fmt.Println(rest)
		fmt.Println(flag)

		if flag {
			zArr[rest] = append(zArr[rest], s[i])
		}else {
			zArr[numRows-rest-1] = append(zArr[numRows-rest-1], s[i])
		}

		if ( rest == numRows-2 && i != 0 ) || (i == 0 && numRows == 2) {
			flag = !flag
		}

	}

	for i:=0; i<numRows; i++ {
		for j:=0; j<len(zArr[i]); j++ {
			result = append(result, zArr[i][j])
		}
	}

	for i:=0; i<numRows; i++ {
		fmt.Println(zArr[i])
	}

	var resStr string = string(result[:])

	return resStr
}