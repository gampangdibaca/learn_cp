package main

import "fmt"

func main() {
	fmt.Println(taumBday(10, 10, 1, 1, 1))
	fmt.Println(taumBday(5, 9, 2, 3, 4))
	fmt.Println(taumBday(3, 6, 9, 1, 1))
	fmt.Println(taumBday(7, 7, 4, 2, 1))
	fmt.Println(taumBday(3, 3, 9, 2, 1))
}

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// Write your code here
	if bc > z+wc {
		return int64(wc) * int64(b+w) + (int64(b) * int64(z))
	} else if wc > z+bc {
		return int64(bc) * int64(b+w) + (int64(w) * int64(z))
	} else {
		return (int64(b)*int64(bc))+(int64(w)*int64(wc))
	}
}

