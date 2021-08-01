package main

import (
	"log"
)

func main() {
	cards := newDeck()
	err := cards.saveToFile("myFile")
	if err != nil {
		log.Fatal(err)
	}
}
