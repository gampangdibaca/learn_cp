package main

import "fmt"

func main() {
	fmt.Println(equalizeArray([]int32{37, 32, 97, 35, 76, 62}))
}

func equalizeArray(arr []int32) int32 {
	// Write your code here
	hm := make(map[int32]int32)
	highest:=int32(1)
	for _, v := range arr {
		_, ok := hm[v]
		if !ok {
			hm[v]=1
		} else {
			hm[v]++
			if highest < hm[v] {
				highest = hm[v]
			}
		}
	}
	return int32(len(arr))-highest
}