package main

import (
	"fmt"
	"iamargus95/Gophercises/BlackJack/deck"
	"strings"
)

type Hand []deck.Card

func main() {

	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}

	var input string
	for input != "s" {
		fmt.Println("Player :", player)
		fmt.Println("Dealer :", dealer.DealerString())
		fmt.Println("What do you wanna do? (h)it or (s)tand? :")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}
	fmt.Println("----- FINAL HANDS -----")
	fmt.Println("Player: ", player)
	fmt.Println("Dealer: ", dealer)

}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func (h Hand) String() string {
	str := make([]string, len(h))
	for i := range h {
		str[i] = h[i].String()
	}
	return strings.Join(str, ", ")
}
