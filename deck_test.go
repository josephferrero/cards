package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLenth := 48

	if len(d) != expectedLenth {
		t.Errorf("Expected deck length of 48, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, got %v", d[0])
	}

	if d[47] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, got %v", d[50])
	}
}

func TestReadDeck(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	nd := newDeckFromFile("_decktesting")

	if len(nd) != 48 {
		t.Errorf("Exptected 16 cards in deck got %v", len(nd))
	}

	os.Remove("_decktesting")
}
