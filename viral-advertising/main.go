package main

import "fmt"

func main() {
	fmt.Println(4)
	fmt.Println(5)
}

func viralAdvertising(n int32) int32 {
	// Write your code here
	cumulative:=int32(0)
	peoples:=int32(5)
	for i:=int32(0); i < n; i++{
		cumulative+=peoples/2
		peoples=peoples/2*3
	}
	return cumulative
}