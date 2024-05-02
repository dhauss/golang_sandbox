package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	numSlice := []int{4, 5, 6, 2, 7, 8}

	resString := evenOrOdd(numSlice)

	fmt.Println(resString)

}

func evenOrOdd(intSlice []int) string {
	var buffer bytes.Buffer

	for _, num := range intSlice {
		if num%2 == 0 {
			buffer.WriteString(strconv.Itoa(num))
			buffer.WriteString(" is even\n")
		} else {
			buffer.WriteString(strconv.Itoa(num))
			buffer.WriteString(" is odd\n")
		}
	}

	return buffer.String()
}
