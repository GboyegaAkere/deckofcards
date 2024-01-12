package main

import "fmt"

// func main() {
// 	// var cards string = "Ace of Spades"
// 	cards := "Ace of Spades"
// 	cards = "five of Diamonds"
// 	fmt.Println(cards)
// }

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
