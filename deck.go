package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// this (d deck) is a receiver! 
func (d deck) print() {
	for i, card := range d { // we use the := here, because we are initializing a new one on each iter.
		fmt.Println(i, card)
	}
}