package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("testCards.txt")
	cards := newDeckFromFile("./testCards.txt")
	cards.print()
}
