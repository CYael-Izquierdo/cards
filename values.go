package main

type Value int8

const (
	Ace = iota << 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Queen
	King
)

func (v Value) String() string {
	switch v {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return ""
	}
}

var values = []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Queen, King}
