package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected the deck to have 52 cards but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, got: %v", d[0])
	}

	lastCard := d[len(d)-1]

	if lastCard != "Joker of Clubs" {
		t.Errorf("Expected last card to be Joker of Clubs, got: %v", lastCard)
	}
}
