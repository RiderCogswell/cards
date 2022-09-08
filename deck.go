package main

import (
	"fmt"
	"os"
	"strings"
)

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

// this annotation declares that we will be returning two values of type deck
func deal(d deck, handSize int) (deck, deck) { 
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string)  {
	os.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is a permission code that allows anyone to read or write
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		 fmt.Println("Error:", err)
		 os.Exit(1)
	}
}