package main

import "fmt"

// A special type to represent card Values
type Value int

/*
	iota will set
		Two = 0
		Three = 1
		Four = 2
	and so on
*/
const (
	Two Value = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (v Value) String() string {
	s := ""
	switch v {
		case Two: s = "2"
		case Three: s = "3"
		case Four: s = "4"
		case Five: s = "5"
		case Six: s = "6"
		case Seven: s = "7"
		case Eight: s = "8"
		case Nine: s = "9"
		case Ten: s = "10"
		case Jack: s = "Jack"
		case Queen: s = "Queen"
		case King: s = "King"
		case Ace: s = "Ace"
	}
	return s
}

type Suit int

const (
	Clubs Suit = iota
	Hearts
	Diamonds
	Spades
)

func (u Suit) String() string {
	s := ""
	switch u {
		case Clubs: s = "Clubs"
		case Hearts: s = "Hearts"
		case Diamonds: s = "Diamonds"
		case Spades: s = "Spades"
	}
	return s
}

func main() {
	vs := []Value{
		Two, Three, Four, 
		Five, Six, Seven, 
		Eight, Nine, Ten, 
		Jack, Queen, King, Ace}
	for _,v := range vs {
		fmt.Println(v)
	}
	ss := []Suit{Clubs, Hearts, Diamonds, Spades}
	for _,v := range ss {
		fmt.Println(v)
	}
}