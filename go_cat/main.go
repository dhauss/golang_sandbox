package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		bs, err := os.ReadFile(arg)

		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Printf("%v", string(bs))
	}
}
