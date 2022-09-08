package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5) // this sets two new vars of hand, and remaining, both with types of deck

	hand.print()
	remainingCards.print()
}