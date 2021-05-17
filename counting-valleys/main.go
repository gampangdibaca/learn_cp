package main

import "fmt"

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU"))
	fmt.Println(countingValleys(12, "DDUUDDUDUUUD"))
}

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	prev := int32(0)
	curr := int32(0)
	valley := int32(0)
	for _, v := range path {
		prev = curr
		if v == 'U' {
			curr = prev + 1
		} else {
			curr = prev - 1
		}
		if prev == -1 && curr == 0 {
			valley += 1
		}
	}
	return valley
}