package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	t.Run("a new deck creates 52 new cards", func(t *testing.T) {
		got := len(deck)
		want := 52

		if got != want {
			t.Errorf("Expected a deck of %v cards, Got a deck of %v cards", got, want)
		}
	})

	t.Run("the first card is always an Ace of Spades", func(t *testing.T) {
		got := deck[0]
		want := "Ace of Spades"

		if got != want {
			t.Errorf("Expected a card of %v, Got a card of %v", got, want)
		}
	})

	t.Run("the last card is always a King of Clubs", func(t *testing.T) {
		got := deck[len(deck)-1]
		want := "King of Clubs"

		if got != want {
			t.Errorf("Expected a card of %v, Got a card of %v", got, want)
		}
	})
}

func TestSaveToDeckAndReadDeckFromFile(t *testing.T) {
	defer os.Remove("_cardtest")

	t.Run("a deck can be saved to a file and a new deck can be created from a file", func(t *testing.T) {
		deck := newDeck()
		err1 := deck.saveToFile("_cardtest")

		if err1 != nil {
			t.Error(err1)
		}

		deckFromFile, err2 := getDeckFromFile("_cardtest")

		if err2 != nil {
			t.Error(err2)
		}

		if len(deckFromFile) != 52 {
			t.Errorf("Expected a deck of 52 cards, Got a deck of %v cards", len(deckFromFile))
		}
	})
}
