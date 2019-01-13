package main

import "fmt"

func main() {

	// only use := when defining a new variable
	// can just assign new values to card after the first :=
	//  ex: card = "Five of Diamonds"

	// cards := []string{newCard(), newCard()}
	cards := newDeck()
	fmt.Println(cards.toString())

}
