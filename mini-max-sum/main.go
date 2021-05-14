package main

import "fmt"

func main() {
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
}

func miniMaxSum(arr []int32) {
	// Write your code here
	min:=0
	max:=0
	sum:=int64(0)
	for i, v := range arr {
		if arr[min] > v {
			min = i
		}
		if arr[max] < v {
			max = i
		}
		sum+=int64(v)
	}
	fmt.Printf("%v %v", sum-int64(arr[max]), sum-int64(arr[min]))
}
