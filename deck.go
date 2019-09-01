package main

import "math/rand"

type Deck struct {
	Cards []Card
}

// adapted from https://gist.github.com/montanaflynn/4cc2779d2e353d7524a7bdce57869a75
// Shuffle the deck
func (deck Deck) Shuffle() Deck {
	for i := 1; i < len(deck.Cards); i++ {
		// Create a random int up to the number of cards
		r := rand.Intn(i + 1)

		// If the the current card doesn't match the random  int we generated then we'll switch them out
		if i != r {
			deck.Cards[r], deck.Cards[i] = deck.Cards[i], deck.Cards[r]
		}
	}

	return deck
}

// adapted from https://gist.github.com/montanaflynn/4cc2779d2e353d7524a7bdce57869a75
func (deck *Deck) Deal() Card {
	card := deck.Cards[0]

	deck.Cards = deck.Cards[1:] // remove card from deck

	return card
}
