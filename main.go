package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// assumptions:
//   Type: ["2", "3, "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"]
//   Suit: ["Diamonds", "Clubs", "Hearts", "Spades"]
type Card struct {
	Type string
	Suit string
}

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

type Deck struct {
	Cards []Card
}

// adapted from https://gist.github.com/montanaflynn/4cc2779d2e353d7524a7bdce57869a75
// Shuffle the deck
func (d Deck) Shuffle() Deck {
	for i := 1; i < len(d.Cards); i++ {
		// Create a random int up to the number of cards
		r := rand.Intn(i + 1)

		// If the the current card doesn't match the random  int we generated then we'll switch them out
		if i != r {
			d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
		}
	}
	return d
}

// adapted from https://gist.github.com/montanaflynn/4cc2779d2e353d7524a7bdce57869a75
func (d *Deck) Deal() Card {
	card := d.Cards[0]

	d.Cards = d.Cards[1:] // remove card from deck

	return card
}

// adapted from https://gist.github.com/montanaflynn/4cc2779d2e353d7524a7bdce57869a75
// New creates a deck of cards to be used
func NewDeck() (deck Deck) {

	types := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

	suits := []string{"Diamonds", "Clubs", "Hearts", "Spades"}

	// Loop over each type and suit appending to the deck
	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type: types[i],
				Suit: suits[n],
			}
			deck.Cards = append(deck.Cards, card)
		}
	}
	return
}

// Seed our randomness with the current time
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	deck := NewDeck()
	deck.Shuffle()

	for i := 0; i < 2; i++ {
		card := deck.Deal()
		fmt.Printf("Card: %s of %s, value: %d. ", card.Type, card.Suit, card.IntValue())
		fmt.Printf("# Cards remaining: %d", len(deck.Cards))
		fmt.Println()
	}
}
