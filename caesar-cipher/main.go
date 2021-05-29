package main

import "fmt"
//https://www.hackerrank.com/challenges/caesar-cipher-1/problem
func main() {
	fmt.Println(caesarCipher("middle-Outz", 2))
	fmt.Println(caesarCipher("Always-Look-on-the-Bright-Side-of-Life", 5))
}

func caesarCipher(s string, k int32) string {
	temp:=""
	k = k%26
	for _, v := range s {
		if int(v) >= 97 && int(v) <= 122{
			if int(v) + int(k) > 122{
				temp += string(rune(96+(int(v) + int(k) - 122)))
			} else {
				temp += string(rune(int(v)+int(k)))
			}
		} else if int(v) >= 65 && int(v) <= 90{
			if int(v) + int(k) > 90{
				temp += string(rune(64+(int(v) + int(k) - 90)))
			} else {
				temp += string(rune(int(v)+int(k)))
			}
		} else {
			temp+=string(v)
		}

	}
	return temp
}