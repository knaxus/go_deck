package main

import (
	"os"
)

func main() {
	// give name to the file
	fileName := "deckInFile"

	// removeFile if already exists
	os.Remove(fileName)

	// create a deck of cads
	cards := newDeck()

	// save the deck in the file
	cards.saveToFile(fileName)

	// load the cards from file
	cards = newDeckFromFile("./" + fileName)

	// shuffle the deck
	cards.shuffle()

	// print the cards
	cards.print()
}
