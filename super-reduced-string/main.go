package main

import "fmt"

func main() {
	fmt.Println(superReducedString("aaabccddd"))
	fmt.Println(superReducedString("aa"))
	fmt.Println(superReducedString("baab"))
}

func superReducedString(s string) string {
	// Write your code here
	temp:=s
	for i:=0; ;i++{
		if i < 0 {
			i = 0
		}
		if temp=="" || i == len(temp)-1{
			break
		}
		if temp != "" && i != len(temp)-1{
			if temp[i] == temp[i+1] {
				temp = temp[:i] + temp[i+2:]
				i -= 2
			}
		}
	}
	if len(temp) == 0 {
		return "Empty String"
	}
	return temp
}