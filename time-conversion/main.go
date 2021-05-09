package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(timeConversion("11:50:00PM"))

	a := "1"
	b := "1"
	fmt.Println(a[len(a)-len(b):])
	fmt.Println("---------")
	fmt.Println(a[:len(a)-len(b)])
}

func timeConversion(s string) string {
	// Write your code here
	finalString := ""
	hour, err := strconv.ParseInt(s[:2], 0, 8)
	checkError(err)
	if s[8:] == "AM"{
		if hour == 12 {
			finalString = finalString + "00" + s[2:8]
		} else {
			finalString = s[:8]
		}
		return finalString
	}
	if hour < 12 {
		finalString = strconv.Itoa(int(hour)+12) + s[2:8]
	} else {
		finalString = s[:8]
	}
	return finalString
}


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}