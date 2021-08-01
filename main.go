package main

func main() {
	cards := newDeck()

	hand, _ := deal(cards, 6)
	hand.shuffle()
	hand.print()
}
