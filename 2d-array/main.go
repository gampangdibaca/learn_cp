package main

import "fmt"

func main() {
	//arr := [][]int32{
	//	{1,1,1,0,0,0},
	//	{0,1,0,0,0,0},
	//	{1,1,1,0,0,0},
	//	{0,0,2,4,4,0},
	//	{0,0,0,2,0,0},
	//	{0,0,1,2,4,0},
	//}
	arr := [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}
	fmt.Println(hourglassSum(arr))
}

func hourglassSum(arr [][]int32) int32 {
	// Write your code here
	highestSum := int32(-1000)
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			sum := arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] + arr[i][j] + arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]
			if sum > highestSum {
				highestSum = sum
			}
		}
	}
	return highestSum
}