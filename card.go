package main

import (
	"strconv"
	"strings"
)

// assumptions:
//   Type: ["2", "3, "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"]
//   Suit: ["Diamonds", "Clubs", "Hearts", "Spades"]
type Card struct {
	Type string
	Suit string
}

// assumptions: ace value is always high
func (c Card) IntValue() int {
	switch lowerType := strings.ToLower(c.Type); lowerType {
	case "jack":
		return 11
	case "queen":
		return 12
	case "king":
		return 13
	case "ace":
		return 14
	default:
		result, _ := strconv.Atoi(lowerType)
		return result
	}
}
