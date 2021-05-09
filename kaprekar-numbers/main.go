package main

import (
	"fmt"
	"strconv"
)

func main() {
	kaprekarNumbers(1, 99999)
}

func kaprekarNumbers(p int32, q int32) {
	// Write your code here
	isKaprekarFound := false
	left := int64(0)
	for i := int64(p); i <= int64(q); i ++ {
		sqrt := i*i
		strI := fmt.Sprintf("%v", i)
		strSqrt := fmt.Sprintf("%v", sqrt)
		strR := strSqrt[len(strSqrt)-len(strI):]
		strRlen := len(strR)
		for j := 0; j < strRlen; j++ {
			if strR[0] != '0'{
				break
			}
			if len(strR) > 1 {
				strR = strR[1:]
			}
		}
		right, err := strconv.ParseInt(strR, 0, 32)
		checkError(err)
		if len(strSqrt[:len(strSqrt)-len(strI)]) > 0 {
			left, err = strconv.ParseInt(strSqrt[:len(strSqrt)-len(strI)], 0, 32)
			checkError(err)
		} else {
			left = int64(0)
		}
		if right + left == int64(i) {
			fmt.Printf("%v ", i)
			isKaprekarFound = true
		}
	}
	if isKaprekarFound == false {
		fmt.Println("INVALID RANGE")
	}
}


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}