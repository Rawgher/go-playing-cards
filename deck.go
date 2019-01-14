package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type of 'deck'
// which is a slice of strings

type deck []string

// this extends deck to make it a slice of string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// use _ for an index when you need to use a for loop but don't need to use that value
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// append will add the card to the slice, making a new slice from the first one declared
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	// (d deck) is a reciever of type deck which gives anything inside it the ability to call the print method
	// recievers are usually one or two letters that reference the type you made, but it can be whatever you want as long as you reference it properly

	// how to iterate over the slice cards
	// for loop that remakes card every time it iterates
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// array is a fixed length list of values
// slices are arrays that can grow and shrink
// all things in arrays or slices must be the same type (string, int, bool, float64)
// [] = slice

// allows us to use type deck and int in the function
// (deck, deck) says we are expecting to return two values and both are type deck
func deal(d deck, handSize int) (deck, deck) {
	// first will return cards up to handSize, second will do the rest of the cards
	return d[:handSize], d[handSize:]
}

// function to change a deck to a string to then change it to a []byte
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ", ")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
