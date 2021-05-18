package main

import "fmt"

func main() {
	fmt.Println(getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10))
	fmt.Println(getMoneySpent([]int32{4}, []int32{5}, 5))
}

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	highest:=int32(0)
	for _, v := range keyboards {
		for _, w := range drives {
			if v + w <= b && v + w > highest {
				highest = v+w
			}
		}
	}
	if highest == 0 {
		return -1
	}
	return highest
}
