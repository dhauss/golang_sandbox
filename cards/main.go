package main

import "fmt"

var deckSize int = 52

func main() {
	/*	var card string = "Ace of Spades"
		card2 := "Ace of Diamonds"
		card2 = "Five of Diamonds"
		fmt.Println(card, card2)

		deckSize = 52
		fmt.Println(deckSize)

		card3 := newCard()
		fmt.Println(card3)
	*/
	/*
		cards := []string{card, card2, card3}
		fmt.Println(cards)
		cards = append(cards, "Six of Spades")
		fmt.Println(cards)

		for i, card := range cards {
			fmt.Println(i, ":", card)
		}

		for i := range cards {
			fmt.Println(i, ":", cards[i])
		}
	*/
	//myDeck := deck{card, card2, card3}
	//myDeck.print()

	cards := newDeck()

	hand, cards := deal(cards, 5)
	fmt.Println("Cards:")
	hand.print()
	fmt.Println("Deck:")
	cards.print()
}

func newCard() string {
	return "Seven of Diamonds"
}
