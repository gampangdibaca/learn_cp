package main

import "fmt"

func main() {
	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	countA:=int32(0)
	countO:=int32(0)
	for _, v := range apples {
		if a + v >= s && a + v <= t {
			countA++
		}
	}
	for _, v := range oranges {
		if b + v >= s && b + v <= t {
			countO++
		}
	}
	fmt.Println(countA)
	fmt.Println(countO)
}
