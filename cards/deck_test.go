package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected length of 20, got length %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card should be Ace of Spades, not %v", d[0])
	}

	if d[len(d)-1] != "Five of Hearts" {
		t.Errorf("Last card should be Five of Hearts, not %v", d[len(d)-1])

	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
