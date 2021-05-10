package main

import "fmt"

func main() {
	a := []int32{1,2,3,4,5}
	d := int32(4)
	fmt.Println(rotLeft(a, d))
}


func rotLeft(a []int32, d int32) []int32 {
	// Write your code here
	if int32(len(a)) == d {
		return a
	}
	finalArray := make([]int32, len(a))
	copy(finalArray, a)
	for i, v := range a {
		 newIdx := int32(i) - d
		 if newIdx < 0 {
		 	newIdx+=int32(len(a))
		 }
		 finalArray[newIdx] = v
	}
	return finalArray
}