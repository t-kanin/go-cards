package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// create type deck
// which is slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)

		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) toString() string {
	// this will return slices of string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	byteString := []byte(d.toString())
	// 0666 means anyone can read and write this file
	return ioutil.WriteFile(filename, byteString, 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
