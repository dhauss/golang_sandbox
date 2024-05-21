package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	/*
		out, err := os.Create("body.txt")

		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		io.Copy(out, resp.Body)
		out.Close()
	*/

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	//bodyBytes, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(bodyBytes))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes: ", len(bs))

	return len(bs), nil
}
