package main

import "fmt"

func main() {
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33}))
}

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	for i, v := range grades {
		if v < 38 {
			continue
		}
		if v%5 > 2 {
			grades[i] = v + 5 - v%5
		}
	}
	return grades
}
