package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of Deck which is a slice of strings
type Deck []string

func newDeck() Deck {
	deck := Deck{}

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, value.String()+" of "+suit.String())
		}
	}

	return deck
}

func (d Deck) Print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) ToString() string {
	return strings.Join(d, ",")
}

func (d Deck) SaveToFile(fileName string) error {
	return os.WriteFile("./saves/"+fileName, []byte(d.ToString()), 0666)
}

func loadDeckFromFile(fileName string) (Deck, error) {
	bs, err := os.ReadFile("./saves/" + fileName)
	if err != nil {
		return nil, err
	}

	s := strings.Split(string(bs), ",")
	return Deck(s), nil
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
