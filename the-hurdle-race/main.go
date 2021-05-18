package main

import "fmt"

func main() {
	fmt.Println(hurdleRace(4, []int32{1, 6, 3, 5, 2}))
	fmt.Println(hurdleRace(7, []int32{2, 5, 4, 5, 2}))
}

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	potionsNeeded := int32(0)
	for _, v := range height{
		if v > k && v - k > potionsNeeded {
			potionsNeeded = v-k
		}
	}
	return potionsNeeded
}
