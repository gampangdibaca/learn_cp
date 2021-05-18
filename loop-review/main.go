package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	count, _ := strconv.ParseInt(strings.TrimRight(input,"\r\n"), 10, 64)
	for i:=int64(0); i < count; i++ {
		input, _ = reader.ReadString('\n')
		strO:=""
		strE:=""
		for i, v := range strings.TrimRight(input,"\r\n") {
			if i%2==0 {
				strE+=string(v)
			} else {
				strO+=string(v)
			}
		}
		fmt.Printf("%v %v\n", strE, strO)
	}
}