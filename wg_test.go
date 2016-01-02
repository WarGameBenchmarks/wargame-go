package main

import (
	"testing"
	"fmt"
)

func TestCardEquality(t *testing.T) {
	table := [][]Card{
		[]Card{Card{Hearts, Two}, Card{Hearts, Two}},
		[]Card{Card{Spades, Five}, Card{Spades, Five}},
		[]Card{Card{Hearts, Jack}, Card{Clubs, Jack}},
		[]Card{Card{Diamonds, Ace}, Card{Clubs, Ace}},
	}
	
	for _,v := range table {
		a := v[0]
		b := v[1]
		if a.Compare(b) != 0 {
			t.Error("Cards are not equal")
		}		
	}
}

func TestCardNotEqual(t *testing.T) {
	table := [][]Card{
		[]Card{Card{Hearts, Six}, Card{Hearts, Two}},
		[]Card{Card{Spades, Five}, Card{Spades, Eight}},
		[]Card{Card{Clubs, Queen}, Card{Clubs, Jack}},
		[]Card{Card{Diamonds, Ace}, Card{Diamonds, King}},
	}
	
	for _,v := range table {
		a := v[0]
		b := v[1]
		if a.Compare(b) == 0 {
			t.Error("Cards are equal")
		}		
	}	
}

func TestDeckFresh(t *testing.T) {
	deck := Deck{}
	deck.Fresh()

	if len(deck.cards) != 52 {
		t.Error(fmt.Sprintf("Deck contains %d cards", len(deck.cards)))
	}
}

func TestDeckSplit(t *testing.T) {
	deck := Deck{}
	deck.Fresh()

	p1, p2 := deck.Split()

	if len(p1.cards) != 26 || len(p2.cards) != 26 {
		t.Error(fmt.Sprintf("Split decks contains %d and %d cards", len(p1.cards), len(p2.cards)))
	}
}

/*
	Testing Deck.shuffle is untenable. Skipped for now.
*/

func TestDeckGetCard(t *testing.T) {
	deck := Deck{}
	deck.Fresh()

	p1, p2 := deck.Split()

	c1, c2 := p1.GetCard(), p2.GetCard()

	// to quiet the compiler for not using these in the test
	if true {
		c1, c2 = c2, c1
	}

	if len(p1.cards) != 26 || len(p2.cards) != 26 {
		t.Error(fmt.Sprintf("Split decks contains %d and %d cards", len(p1.cards), len(p2.cards)))
	}
}

func TestDeckGiveCard(t *testing.T) {
	deck := Deck{}
	deck.Fresh()
	p1, p2 := deck.Split()

	// notice that &p2 implies that this is a pointer
	p1.GiveCard(&p2)
	if len(p1.cards) != 25 || len(p2.cards) != 27 {
		t.Error(fmt.Sprintf("Split decks contains %d and %d cards", len(p1.cards), len(p2.cards)))
	}
}

func TestDeckGiveCards(t *testing.T) {
	deck := Deck{}
	deck.Fresh()
	p1, p2 := deck.Split()
	
	p1.GiveCards(&p2)
	if len(p1.cards) != 0 || len(p2.cards) != 52 {
		t.Error(fmt.Sprintf("Split decks contains %d and %d cards", len(p1.cards), len(p2.cards)))
	}
}

