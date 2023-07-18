package main

// import "fmt"

func main() {
	cards := newDeck()
	// cards := newDeckFromFile("my_cards.txt")
	cards.shuffle()
	cards.print()
}
