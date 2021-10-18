package blackjack

var cardVal = map[string]int{"ace": 11, "two": 2, "three": 3, "four": 4,
	"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10,
	"jack": 10, "queen": 10, "king": 10}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cardVal[card]
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return cardVal[card1]+cardVal[card2] == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {

	var choice string
	switch {
	case isBlackjack && dealerScore < 10:
		choice = "W"
	case isBlackjack && dealerScore >= 10:
		choice = "S"
	default:
		choice = "P"
	}

	return choice
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	var choice string
	switch {
	case handScore <= 11 || handScore >= 12 && handScore <= 16 && dealerScore >= 7:
		choice = "H"
	default:
		choice = "S"
	}

	return choice
}
