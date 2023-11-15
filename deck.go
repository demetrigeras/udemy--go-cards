package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	CardValues := []string{"Ace", "Two", "Three", "Four"}

	for i, suit := range CardSuits {
		for j, value := range CardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
