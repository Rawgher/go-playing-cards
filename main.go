package main

func main() {

	// only use := when defining a new variable
	// can just assign new values to card after the first :=
	//  ex: card = "Five of Diamonds"

	// cards := []string{newCard(), newCard()}
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

}
