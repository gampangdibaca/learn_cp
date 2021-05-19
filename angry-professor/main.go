package main

import "fmt"

func main() {
	fmt.Println(angryProfessor(3, []int32{-1, -3, 4, 2}))
	fmt.Println(angryProfessor(2, []int32{0, -1, 2, 1}))
}

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	count:=int32(0)
	for _, v := range a {
		if v <=0 {
			count++
			if count >=k{
				return "NO"
			}
		}
	}
	return "YES"
}