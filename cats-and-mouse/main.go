package main

import "fmt"

func main() {
	fmt.Println(catAndMouse(1, 2, 3))
	fmt.Println(catAndMouse(1, 3, 2))
}

func catAndMouse(x int32, y int32, z int32) string {
	ACatD:=int32(0)
	BCatD:=int32(0)

	if x > z {
		ACatD = x - z
	} else {
		ACatD = z - x
	}
	if y > z {
		BCatD = y - z
	} else {
		BCatD = z - y
	}
	if ACatD > BCatD {
		return "Cat B"
	} else if ACatD == BCatD {
		return "Mouse C"
	} else {
		return "Cat A"
	}
}