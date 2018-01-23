package main

import (
	"fmt"
)

func main() {
	cards := []string{"Ace of Diamond", newCard()}
	cards = append(cards, "Six of Spade")

	for i, cards := range cards {
		fmt.Println(i, cards)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
