package main

import "fmt"

func main() {
	fmt.Println(strangeCounter(4))
	fmt.Println(strangeCounter(17))
	fmt.Println(strangeCounter(5555))
}

func strangeCounter(t int64) int64 {
	// Write your code here
	n:=int64(3)
	for {
		if t <= n {
			return n + 1 - t
		}
		t-=n
		n*=2
	}
}