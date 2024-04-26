package main

import "fmt"

var deckSize int = 52

func main() {
	var card string = "Ace of Spades"
	card2 := "Ace of Diamonds"
	card2 = "Five of Diamonds"
	/*
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
	myDeck := deck{card, card2}
	myDeck.toString()

	cards := newDeck()

	hand, cards := deal(cards, 5)
	fmt.Println("Cards:", hand.toString())

	hand.shuffle()
	fmt.Println("Shuffled hand: ", hand.toString())

	//cards.saveToFile("myCards")

	//fmt.Println("\ncardsFromFile")
	//cardsFromFile := newDeckFromFile("myCards")
	//fmt.Println(cardsFromFile.toString())
}

func newCard() string {
	return "Seven of Diamonds"
}
