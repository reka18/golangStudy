package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	L := len(deck.cards)
	if L != 52 {
		t.Errorf("Expected length is 52: %v", L)
	}

	pos1 := deck.cards[0]
	if pos1.suit != "Spades" || pos1.value != "Ace" {
		t.Errorf("Expected Ace of Spades but got %v of %v", pos1.value, pos1.suit)
	}
}
