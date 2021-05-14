package main

import "fmt"

func main() {
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
}

func miniMaxSum(arr []int32) {
	// Write your code here
	min:=0
	max:=0
	for i, v := range arr {
		if arr[min] > v {
			min = i
		}
		if arr[max] < v {
			max = i
		}
	}
	minSum:=int64(0)
	maxSum:=int64(0)
	for i, v := range arr {
		if i != min {
			maxSum+=int64(v)
		}
		if i != max {
			minSum+=int64(v)
		}
	}
	fmt.Printf("%v %v", minSum, maxSum)
}
