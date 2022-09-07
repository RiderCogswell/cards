package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard(), newCard()}

	cards = append(cards, "Six of Spades")

	for i, card := range cards { // we use the := here, because we are initializing a new one on each iter.
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}