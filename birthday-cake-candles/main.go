package main

import "fmt"

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
}

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	count:=int32(0)
	highest:=candles[0]
	for _, v := range candles{
		if highest < v {
			highest = v
			count = 1
		} else if highest == v {
			count++
		}
	}
	return count
}
