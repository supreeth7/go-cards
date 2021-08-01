package main

import (
	"fmt"
	"strings"
)

// Create a new deck type
type deck []string

// Creates a set of cards and returns them as a deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, value := range cardValues {
		for _, suite := range cardSuits {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

// Prints all the cards in the deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Returns two slices of deck with the given hand size
// Return Values - deck 1: deck in hand of size handSize, deck 2: remaining cards in the deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Converts a deck into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
