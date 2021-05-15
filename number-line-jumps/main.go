package main

import "fmt"

func main() {
	fmt.Println(kangaroo(0, 3, 4, 2))
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	if v1 <= v2 || x1+v1 > x2+v2 {
		return "NO"
	}
	for {
		if x1 < x2 {
			x1+=v1
			x2+=v2
		} else if x1 == x2 {
			return "YES"
		} else if x1 > x2 {
			return "NO"
		}
	}
}
