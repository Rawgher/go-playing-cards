package main

func main() {

	// only use := when defining a new variable
	// can just assign new values to card after the first :=
	//  ex: card = "Five of Diamonds"

	// cards := []string{newCard(), newCard()}
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
