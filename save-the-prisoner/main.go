package main

import "fmt"

func main() {
	fmt.Println(saveThePrisoner(559033664, 822024089, 131746755))
	fmt.Println(saveThePrisoner(144, 797951344, 1))
	fmt.Println(saveThePrisoner(999999998, 999999997, 1))
	fmt.Println(saveThePrisoner(999999999, 999969318, 999999999))
	fmt.Println(saveThePrisoner(3, 7, 3))


}

func saveThePrisoner(n int32, m int32, s int32) int32 {
	// Write your code here
	return (m-1+s-1) % n+1
}