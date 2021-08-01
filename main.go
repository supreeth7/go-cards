package main

func main() {
	cards := newDeck()
	hand := cards.deal(3)
	hand.print()

}
