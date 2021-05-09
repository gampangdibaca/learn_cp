package main

import "fmt"

func main() {
	arr := []int32{0, 0, 0, 1, 0, 0}
	fmt.Println(jumpingOnClouds(arr))
}

func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	step:= int32(0)
	for i := 0; i < len(c); i++{
		if i+2 < len(c) && c[i+2] != int32(1) {
			i+=1
		}
		if i < len(c)-1 {
			step += 1
		}
	}
	return step
}