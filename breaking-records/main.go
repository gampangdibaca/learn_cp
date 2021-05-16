package main

import "fmt"

func main() {
	fmt.Println(breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}))
}

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	lowest:= scores[0]
	countL:=int32(0)
	highest:= scores[0]
	countH:=int32(0)
	for _, v := range scores {
		if lowest > v {
			lowest = v
			countL++
		}
		if highest < v {
			highest = v
			countH++
		}
	}
	return []int32{countH, countL}
}