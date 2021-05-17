package main

import "fmt"

func main() {
	fmt.Println(sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}))
}

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	var hm = make(map[int32]int32)
	for _, v := range ar {
		_, ok := hm[v]
		if !ok {
			hm[v] = 1
		} else {
			hm[v]+=1
		}
	}
	count:=0
	for _, v := range hm {
		count+=int(v/2)
	}
	return int32(count)
}