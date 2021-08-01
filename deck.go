package main

import "fmt"

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

func (d deck) deal(n int) deck {
	return d[:n]
}
