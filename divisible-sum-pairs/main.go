package main

import "fmt"

func main() {
	fmt.Println(divisibleSumPairs(3, []int32{1,3,2,6,1,2}))
}

func divisibleSumPairs(k int32, ar []int32) int32 {
	// Write your code here
	count:=int32(0)
	for i:=0; i < len(ar); i++ {
		for j:=i; j < len(ar); j++ {
			if i != j && (ar[i]+ar[j])%k==0{
				count++
			}
		}
	}
	return count
}