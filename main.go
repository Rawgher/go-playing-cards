package main

import "fmt"

func main() {

	// only use := when defining a new variable
	// can just assign new values to card after the first :=
	//  ex: card = "Five of Diamonds"

	// cards := []string{newCard(), newCard()}
	cards := []string{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")
	// append will add the card to the slice, making a new slice from the first one declared

	// how to iterate over the slice cards
	// for loop that remakes card every time it iterates
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}

// array is a fixed length list of values
// slices are arrays that can grow and shrink
// all things in arrays or slices must be the same type (string, int, bool, float64)
// [] = slice
