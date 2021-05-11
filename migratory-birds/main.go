package main

import "fmt"

func main() {
	fmt.Println(migratoryBirds([]int32{1,2,3,4,5,4,3,2,1,3,4}))
}

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	countArray:=[]int{0,0,0,0,0}
	for _, v := range arr {
		countArray[v-1]++
	}
	result:=int32(0)
	resultIndex:=int32(0)
	for i, v := range countArray {
		if result < int32(v) {
			result = int32(v)
			resultIndex = int32(i)
		}
	}
	return resultIndex+1
}