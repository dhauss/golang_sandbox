package main

import "fmt"

// create 'deck' type, slice of strings

type deck []string

func newDeck() deck {
	d := deck{}
	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			d = append(d, val+" of "+suit)
		}
	}

	return d
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
