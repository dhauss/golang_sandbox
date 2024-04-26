package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected length of 20, got length %v", len(d))
	}

}
