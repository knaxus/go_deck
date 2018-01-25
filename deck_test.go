package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected the deck to have 52 cards but got: %v", len(d))
	}
}
