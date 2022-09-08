package main

func main() {
	// use deck type instead
	cards := deck{"Ace of Diamonds", newCard(), newCard()}

	cards = append(cards, "Six of Spades")

	cards.print() // calling the deck print method from deck.go
}

func newCard() string {
	return "Five of Diamonds"
}