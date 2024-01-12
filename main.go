package main

// func main() {
// 	card := newCard()
// 	fmt.Println(card)
// }

// func newCard() string {
// 	return "Five of Diamonds"
// }

// CREATING A SLICE IN GOLANG i.e AN ARRAY IN GOLANG
// func main() {
// 	cards := deck{"Five of Diamond", "Ace of Spades", newCard()}
// 	// ADDING A NEW ELEMENT INTO A SLICE
// 	cards = append(cards, "six of Diamond")

// 	//ITERAING THROUGH THE ARRAY
// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	}
// 	// fmt.Println(cards)
// }

//COMBINING THE TWO FILES TOGETHER
func main() {
	cards := deck{"Five of Diamond", "Ace of Spades", newCard()}
	cards = append(cards, "six of Diamond")

	cards.print()

}

func newCard() string {
	return "Three of Diamond"
}
