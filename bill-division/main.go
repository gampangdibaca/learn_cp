package main

import "fmt"

func main() {
	bonAppetit([]int32{3,10,2,9}, 1, 12)
	bonAppetit([]int32{3,10,2,9}, 1, 7)
}

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	sum:=int32(0)
	for i, v := range bill {
		if int32(i) == k {
			continue
		}
		sum+=v
	}
	if sum / 2 == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - sum/2)
	}
}
