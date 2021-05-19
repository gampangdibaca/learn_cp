package main

import "fmt"

func main() {
	fmt.Println(designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
	"abc"))
	fmt.Println(designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7},
		"zaba"))
}

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	highest:=int32(0)
	for _, v := range word {
		if highest < h[int(v)-97] {
			highest = h[int(v)-97]
		}
	}
	return highest*int32(len(word))
}