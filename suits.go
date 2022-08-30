package main

type Suit int8

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

func (s Suit) String() string {
	switch s {
	case Spades:
		return "Spades"
	case Hearts:
		return "Hearts"
	case Clubs:
		return "Club"
	case Diamonds:
		return "Diamonds"
	default:
		return ""
	}
}

var suits = []Suit{Spades, Hearts, Diamonds, Clubs}
