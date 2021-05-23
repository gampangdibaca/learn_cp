package main

import "fmt"

func main() {
	fmt.Println(squares(3, 9))
	fmt.Println(squares(17, 24))
	fmt.Println(squares(35, 70))
	fmt.Println(squares(100, 1000))
}

func squares(a int32, b int32) int32 {
	// Write your code here
	count:=int32(0)
	for i:=int32(1); ;i++{
		if i*i >= a && i*i <= b {
			count++
		}
		if i*i > b {
			break
		}
	}
	return count
}
