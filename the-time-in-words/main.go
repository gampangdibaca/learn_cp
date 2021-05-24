package main

import "fmt"

func main() {
	fmt.Println(timeInWords(5, 47))
	fmt.Println(timeInWords(5, 27))
	fmt.Println(timeInWords(5, 30))
	fmt.Println(timeInWords(5, 0))
}

func timeInWords(h int32, m int32) string {
	// Write your code here
	hm := map[int32]string {
		1 : "one",
		2 : "two",
		3 : "three",
		4 : "four",
		5 : "five",
		6 : "six",
		7 : "seven",
		8 : "eight",
		9 : "nine",
		10 : "ten",
		11 : "eleven",
		12 : "twelve",
		13 : "thirteen",
		14 : "fourteen",
		16 : "sixteen",
		17 : "seventeen",
		18 : "eighteen",
		19 : "nineteen",
		20 : "twenty",
		21 : "twenty one",
		22 : "twenty two",
		23 : "twenty three",
		24 : "twenty four",
		25 : "twenty five",
		26 : "twenty six",
		27 : "twenty seven",
		28 : "twenty eight",
		29 : "twenty nine",
	}
	if m < 30 {
		if m == 0 {
			return fmt.Sprintf("%s o' clock",hm[h])
		} else if m == 1 {
			return fmt.Sprintf("%s minute past %s", hm[m], hm[h])
		} else if m == 15 {
			return fmt.Sprintf("quarter past %s", hm[h])
		}
		return fmt.Sprintf("%s minutes past %s", hm[m], hm[h])
	} else if m > 30 {
		if 60 - m == 1 {
			return fmt.Sprintf("%s minute to %s", hm[m], hm[h+1])
		} else if 60 - m == 15 {
			return fmt.Sprintf("quarter to %s", hm[h+1])
		}
		return fmt.Sprintf("%s minutes to %s", hm[60-m], hm[h+1])
	} else {
		return fmt.Sprintf("half past %s", hm[h])
	}
}
