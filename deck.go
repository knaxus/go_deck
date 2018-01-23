package main

import (
	"fmt"
)

// create a new type of 'deck'
// which is essentially a slice of stings

type deck []string

// print function to list all the cards in a deck

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
