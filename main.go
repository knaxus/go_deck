package main

func main() {
	cards := newDeck()
	cards.saveToFile("testCards.txt")
}
