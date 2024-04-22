package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	parseCard1 := ParseCard(card1)
	parseCard2 := ParseCard(card2)
	parseCardDealer := ParseCard(dealerCard)
	switch {
	case parseCard1 == 11 && parseCard2 == 11:
		return "P"
	case parseCard1+parseCard2 == 21:
		if parseCardDealer == 11 || parseCardDealer == 10 {
			return "S"
		} else {
			return "W"
		}
	case parseCard1+parseCard2 >= 17 && parseCard1+parseCard2 <= 20:
		return "S"
	case parseCard1+parseCard2 >= 12 && parseCard1+parseCard2 <= 16:
		if parseCardDealer >= 7 {
			return "H"
		} else {
			return "S"
		}
	case parseCard1+parseCard2 <= 11:
		return "H"
	default:
		return ""
	}
}
