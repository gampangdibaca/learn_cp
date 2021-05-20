package main

import "fmt"

func main() {
	fmt.Println(findDigits(2))
	fmt.Println(findDigits(12))
	fmt.Println(findDigits(1012))
}

func findDigits(n int32) int32 {
	// Write your code here
	temp:=n
	count:=int32(0)
	for {
		if temp <= 0 {
			break
		}
		digit:=temp%10
		if digit != 0 && n%digit == 0 {
			count++
		}
		temp/=10
	}
	return count
}