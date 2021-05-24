package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(encryption("have a nice day"))
}

func encryption(s string) string {
	// Write your code here
	temp:=strings.TrimSpace(s)
	// r:= int(math.Floor(float64(len(temp))))
	c:= int(math.Ceil(math.Sqrt(float64(len(temp)))))

	idx:=0
	currRow:=0
	ans:=""
	for i:=0; i < len(temp); i++{
		fmt.Println(idx)
		if idx < len(temp){
			ans+=temp[idx:idx+1]
			if idx+c < len(temp) {
				idx+=c
			} else {
				currRow++
				idx=currRow
				ans+=" "
			}
		}
	}
	return ans
}