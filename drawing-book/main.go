package main

import "fmt"

func main() {
	fmt.Println(pageCount(6, 2))
}

func pageCount(n int32, p int32) int32 {
	// Write your code here
	front := int32(0)
	stepF:=int32(0)
	back := n - n%2
	stepB:=int32(0)
	for {
		if front == p || front+1 == p {
			return stepF
		}
		if back == p || back+1 == p {
			return stepB
		}
		front+=2
		back-=2
		stepF+=1
		stepB+=1
	}
}