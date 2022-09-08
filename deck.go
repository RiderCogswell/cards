package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// we do not need a receiver because we do not need any extra functionality
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// this (d deck) is a receiver! meaning any var
// of type deck can use these receiver functions
func (d deck) print() {
	for i, card := range d { // we use the := here, because we are initializing a new one on each iter.
		fmt.Println(i, card)
	}
}