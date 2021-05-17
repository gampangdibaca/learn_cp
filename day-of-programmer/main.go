package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(dayOfProgrammer(1800))
	fmt.Println(dayOfProgrammer(2500))
	fmt.Println(dayOfProgrammer(2400))
}

func dayOfProgrammer(year int32) string {
	// Write your code here
	if year <= 1917 && year %4 == 0 {
		return "12.09." + strconv.Itoa(int(year))
	} else if year > 1918 && year%400 == 0{
		return "12.09." + strconv.Itoa(int(year))
	} else if year > 1918 && year%4 == 0 && year%100 != 0{
		return "12.09." + strconv.Itoa(int(year))
	} else if year == 1918{
		return "26.09." + strconv.Itoa(int(year))
	} else {
		return "13.09." + strconv.Itoa(int(year))
	}
}
