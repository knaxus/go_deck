package main

import (
	"os"
	"testing"
)

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

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"
	// remove any test file lying around
	os.Remove(testFileName)

	d := newDeck()
	// save to file
	d.saveToFile(testFileName)

	loadedDeck := newDeckFromFile(testFileName)

	// check if the length of the loaded deck is 52
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck to have 52 cards found : %v", len(loadedDeck))
	}

	// delete the test file again
	os.Remove(testFileName)
}
