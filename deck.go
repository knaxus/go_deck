package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// function to convert a deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save a deck to file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// get deck from file
func newDeckFromFile(fileName string) deck {
	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}

	s := strings.Split(string(fileContent), ",")
	return deck(s)
}

// shuffle function
func (d deck) shuffle() deck {
	// create a new source, passing time to get a true random number
	source := rand.NewSource(time.Now().UnixNano())
	// create a random number
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
