package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// we do not need a receiver because we do not need any extra functionality
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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
	for i, card := range d { // we use the := here, because we are initializing a new one on each iteration
		fmt.Println(i, card)
	}
}

// this annotation declares that we will be returning two values of type deck
func deal(d deck, handSize int) (deck, deck) { // WE DO NOT MAKE THIS A RECEIVER BECAUSE IT WOULD BE MODIFYING THE CARDS SLICE
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) {
	os.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is a permission code that allows anyone to read or write
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) // bs = byte slice  
	if err != nil {
		 fmt.Println("Error:", err)
		 os.Exit(1) // exit program completely
	}

	s := strings.Split(string(bs), ",") // turn into string slice
	return deck(s) // return this new deck
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // using the rand.NewSource() std package, and passing in a random time as an int64, we are successfully getting a new randomization each time
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1) // random number between the indexed length of the cards -> d

		d[i], d[newPosition] = d[newPosition], d[i] // one line swap algo 
	}
}