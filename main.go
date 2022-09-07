package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" <- this is the same as below
	card := "Ace of Spades" // only use this syntax when DEFINING a new variable and when we want the type inferred

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}