package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected lenght of deck is 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card was Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Cubs" {
		t.Errorf("Expected first card was Four of Cubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected lenght of deck is 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
