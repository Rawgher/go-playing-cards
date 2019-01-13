package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings

type deck []string

// this extends deck to make it a slice of string

func (d deck) print() {

	// how to iterate over the slice cards
	// for loop that remakes card every time it iterates

	for i, card := range d {
		fmt.Println(i, card)
	}
}
