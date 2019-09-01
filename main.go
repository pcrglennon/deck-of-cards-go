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
		switch i % 4 {
		case 0:
			activityTypes = append(activityTypes, "push_up")
		case 2:
			activityTypes = append(activityTypes, "plank")
		default:
			activityTypes = append(activityTypes, "rest")
		}
	}

	return
}

func promptMessage(cardsRemaining int) string {
	if cardsRemaining%5 == 0 {
		return fmt.Sprintf("\nPress Enter to draw next card (%d cards remaining):", cardsRemaining)
	} else {
		return "\nPress Enter to draw next card:"
	}
}

func outputActivity(card Card, activityType string) {
	fmt.Printf("\nCard: %s of %s ---> ", card.Type, card.Suit)

	normalizedCardValue := normalizeCardValue(card.IntValue())
	activityValue := valueForActivity(normalizedCardValue, activityType)

	switch activityType {
	case "push_up":
		fmt.Printf("%d Push-Ups\n", activityValue)
	case "plank":
		fmt.Printf("%d-second Plank\n", activityValue)
		fmt.Println("Starting in 2 seconds…")
		time.Sleep(2 * time.Second)
		displayActivityTimer(activityValue, 10)
	default:
		fmt.Printf("%d-second Rest\n", activityValue)
		displayActivityTimer(activityValue, 5)
	}
}

// override card-value rules:
//   - treat all face-value cards as 10
//   - treat all Aces as 11
func normalizeCardValue(cardValue int) int {
	switch cardValue {
	case 11, 12, 13:
		return 10
	case 14:
		return 11
	default:
		return cardValue
	}
}

func valueForActivity(cardValue int, activityType string) int {
	switch activityType {
	case "plank":
		return cardValue * 3
	case "rest":
		return cardValue * 2
	default:
		return cardValue
	}
}

func displayActivityTimer(activityTimeSeconds int, intervalSeconds int) {
	fmt.Print("Starting: 0…")

	time.Sleep(1 * time.Second)

	for i := 1; i < activityTimeSeconds; i++ {
		if (activityTimeSeconds - i) < intervalSeconds { // output update for each second of final interval
			fmt.Printf(" %d,", i)
		} else if (i % intervalSeconds) == 0 { // output update for each interval
			fmt.Printf(" %d…", i)
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Printf(" %d!\n", activityTimeSeconds)
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
		card := deck.Deal()
		outputActivity(card, activityTypes[i])

		// wait for Enter input to start next Activity
		fmt.Print(promptMessage(len(deck.Cards)))
		reader.ReadString('\n')
	}

	fmt.Println("\n\n------------ Done! ------------")
}
