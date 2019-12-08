package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeckFromFile("my_cards")
	// hand, remainingCards := deal(cards, 3)
	// hand.print()
	// remainingCards.print()
	cards.print()
	cards.shuffle()
	cards.print()
}
