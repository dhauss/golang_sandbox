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
	}

	c := make(chan string)

	for _, site := range sites {
		go checkStatus(site, c)
	}

	for i := 0; i < 100; i++ {
		go checkStatus(<-c, c)
	}

}

func checkStatus(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Printf("GET error with url: %v\n", site)
		c <- site
	} else {
		fmt.Printf("No problems with url: %v\n", site)
		c <- site
	}
}
