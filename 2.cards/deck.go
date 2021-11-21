package main

import "fmt"

// create a new type of deck
//which is a slice of string
type deck []string

//this says it is equal to strings

// we dont need revciver here as still we dont have any deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

//here inside func we are putting reciver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
