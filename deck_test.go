package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected Length of 20, but got ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of ace of spades but got ", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of four of clubs but got ", d[len(d)-1])
	}
}
