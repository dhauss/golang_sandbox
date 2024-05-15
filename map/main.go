package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
	}

	printMap(colors)

	colors2 := make(map[string]string)
	colors2["white"] = "ffffff"
	fmt.Println(colors2)
	delete(colors2, "white")
	fmt.Println(colors2)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v: %v\n", color, hex)
	}
}
