package main

import (
	"os"
	"strings"
	"testing"
)

const deckName = "_deckTesting"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLength := len(suits) * len(values)

	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %d, but got %d", expectedLength, len(d))
	}
}

func TestDeck_SaveToFile(t *testing.T) {
	os.Remove("./saves/" + deckName)

	d := newDeck()
	d.SaveToFile(deckName)

	rf, err := os.ReadFile("./saves/" + deckName)
	if err != nil {
		t.Errorf("File not created")
	}
	sf := string(rf)

	if sf != strings.Join(d, ",") {
		t.Errorf("Wrong deck format")
	}
}
