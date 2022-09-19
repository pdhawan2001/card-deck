package main

func main() {
	// var card string = "Ace of spades" // only a string will be assigned to this, as we have mentioned string on it
	// card := "Ace of Spades" // equivalent to above line

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
