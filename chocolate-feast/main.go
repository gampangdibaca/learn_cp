package main

import "fmt"

func main() {
	fmt.Println(chocolateFeast(6, 2, 2))
}

func chocolateFeast(n int32, c int32, m int32) int32 {
	// Write your code here
	count:=int32(0)
	chocolateBought := n / c
	count+=chocolateBought
	chocolateWrap:=chocolateBought
	for {
		if chocolateWrap < m {
			return count
		} else {
			chocolateBought = chocolateWrap/m
			chocolateWrap = chocolateWrap%m + chocolateBought
			count+=chocolateBought
		}
	}

}