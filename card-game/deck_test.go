package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	deckSize := 40

	if len(d) != deckSize {
		t.Errorf("Expected deck length of %d, but got, %d", deckSize, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %s", d[0])
	}

	lastCard := d[len(d)-1]
	if lastCard != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs, but got %s", lastCard)
	}
}

func TestReadDeckFromFileAndSaveToFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readDeckFromFile("_decktesting")

	if len(loadedDeck) != 40 {
		t.Errorf("Expected 40 cards in deck, but got %d", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
