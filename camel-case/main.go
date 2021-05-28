package main

import "fmt"

func main() {
	fmt.Println(camelcase("saveChangesInTheEditor"))
}

func camelcase(s string) int32 {
	// Write your code here
	count:=int32(1)
	for _, v := range s {
		if int(v) >= 65 && int(v) <= 90 {
			count++
		}
	}
	return count
}