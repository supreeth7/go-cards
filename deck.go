package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
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

// Randomizes the slice order in a deck of cards
func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := range d {
		swapPos := rand.Intn(len(d) - 1)
		d[i], d[swapPos] = d[swapPos], d[i]
	}
}

// Converts a deck into a continuous comma separated string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save the deck string to a user defined file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0776)
}

// Read string from a file and convert it to a deck
func getDeckFromFile(fileName string) (deck, error) {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	deckString := string(bytes)

	deckSlice := strings.Split(deckString, ",")

	return deck(deckSlice), nil
}
