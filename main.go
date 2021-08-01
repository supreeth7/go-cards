package main

import "log"

func main() {
	cards := newDeck()

	hand, cardsLeft := deal(cards, 6)
	hand.shuffle()
	hand.print()

	fileSaveError := cardsLeft.saveToFile("deck-remains")

	if fileSaveError != nil {
		log.Fatal(fileSaveError)
	}

	retrievedDeck, fileReadError := getDeckFromFile("deck-remains")

	if fileReadError != nil {
		log.Fatal(fileReadError)
	}

	retrievedDeck.shuffle()
	retrievedDeck.print()
}
