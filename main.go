package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// adapted from https://gist.github.com/montanaflynn/4cc2779d2e353d7524a7bdce57869a75
// New creates a deck of cards to be used
func initializeDeck() (deck Deck) {
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

	deck.Shuffle()

	return
}

// [push_up, rest, plank, rest, ...]
func initializeActivityTypes(length int) (activityTypes []string) {
	for i := 0; i < length; i++ {
		activityType := ""

		switch i % 4 {
		case 0:
			activityType = "push_up"
		case 2:
			activityType = "plank"
		default:
			activityType = "rest"
		}

		activityTypes = append(activityTypes, activityType)
	}

	return
}

func outputActivity(card Card, activityType string) {
	cardValue := card.IntValue()
	fmt.Printf("Card: %s of %s ---> ", card.Type, card.Suit)

	switch activityType {
	case "push_up":
		fmt.Printf("%d Push-Ups\n", cardValue)
	case "plank":
		fmt.Printf("%d-second Plank\n", cardValue)
	default:
		fmt.Printf("%d-second Rest\n", cardValue)
	}
}

// Seed our randomness with the current time
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	deck := initializeDeck()
	activityTypes := initializeActivityTypes(len(deck.Cards))

	for i := 0; i < len(deck.Cards); i++ {
		if len(deck.Cards)%5 == 0 {
			fmt.Printf("\n(%d cards remaining)\n\n", len(deck.Cards))
		}

		// wait for input
		fmt.Println("Press Enter to draw next card")
		reader.ReadString('\n')

		card := deck.Deal()
		outputActivity(card, activityTypes[i])
	}

	fmt.Println("\n\n------------ Done! ------------")
}
