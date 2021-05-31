package main

import "fmt"

func main() {
	fmt.Println(marsExploration("SOSSPSSQSSOR"))
	fmt.Println(marsExploration("SOSSOT"))
	fmt.Println(marsExploration("SOSSOSSOS"))
}

func marsExploration(s string) int32 {
	// Write your code here
	count:=int32(0)
	for i, v := range s{
		if i%3 == 1 && v != 'O' {
			count++
		}else if i%3 != 1 && v != 'S'{
			count++
		}
	}
	return count
}