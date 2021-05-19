package main

import "fmt"

func main() {
	fmt.Println(utopianTree(4))
	fmt.Println(utopianTree(3))
	fmt.Println(utopianTree(2))
}

func utopianTree(n int32) int32 {
	// Write your code here
	sum:=int32(0)
	for i:=int32(0); i <= n; i++{
		if i%2==0{
			sum++
		} else {
			sum*=2
		}
	}
	return sum
}