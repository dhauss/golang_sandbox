package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"sdfssd",
	}

	for _, site := range sites {
		checkStatus(site)
	}
}

func checkStatus(site string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Printf("Get error with url: %v\n", site)
	} else {
		fmt.Printf("No problems with url: %v\n", site)
	}
}
