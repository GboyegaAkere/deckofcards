package main

import "fmt"

//CREATING A NEW TYPE OF DECK
//WHICH IS A SLICE OF STRING

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
