package main

import (
	"fmt"
)

// create a new type of 'deck'
// which is essentially a slice of stings

type deck []string

// function to create all the 52 cards of a deck
func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Joker"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// function to split the deck into 2 parts, deal in hand
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// print function to list all the cards in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
