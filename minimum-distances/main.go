package main

import "fmt"

func main() {
	fmt.Println(minimumDistances([]int32{1, 2, 3, 4, 10}))
	fmt.Println(minimumDistances([]int32{7, 1, 3, 4, 1, 7}))
}

func minimumDistances(a []int32) int32 {
	// Write your code here
	hm := make(map[int32]int32)
	min:=int32(-1)
	for i, v := range a {
		_, ok := hm[v]
		if !ok {
			hm[v] = int32(i)
		} else {
			if min == -1 || min > int32(i) - hm[v]{
				min = int32(i) - hm[v]
			}
		}
	}
	return min
}