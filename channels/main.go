package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://34t0efuhnaer.crm",
		"sdfkuuh",
	}

	c := make(chan string)

	for _, site := range sites {
		go checkStatus(site, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkStatus(link, c)
		}(l)
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
