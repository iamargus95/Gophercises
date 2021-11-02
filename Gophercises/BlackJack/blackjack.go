package main

import (
	"fmt"
	deck "iamargus95/Gophercises/Gophercises/DeckofCards"
)

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	for i := 0; i < 10; i++ {
		card, cards = cards[0], cards[1:]
		fmt.Println(card)
	}
}
