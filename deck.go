package main

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck struct {
	cards []card
}

func newDeck() deck {
	d := deck{}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range suits {
		for _, value := range values {
			d.cards = append(d.cards, card{suit, value})
		}
	}
	return d
}

func (d deck) print() {
	for _, card := range d.cards {
		card.print()
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return deck{d.cards[:handSize]}, deck{d.cards[handSize : handSize*2]}
}

func (d deck) toString() string {
	str := ""
	for _, card := range d.cards {
		str = str + card.toString() + ","
	}
	return str[:len(str)-1]
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, d.toByteSlice(), 0666)
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(d.cards),
		func(i int, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		println("ERROR: ", err.Error())
		os.Exit(1)
	}
	deckString := string(byteSlice)
	deckStrings := strings.Split(deckString, ",")
	deck := deck{}
	for _, cString := range deckStrings {
		deck.cards = append(deck.cards, buildCard(cString))
	}
	return deck
}
