package main

import (
	"fmt"
)

func main() {

	empty := Deck{}
	full_deck := empty.fresh()
	full_deck.shuffle()

	p1, p2 := full_deck.split()

	fmt.Println(len(p1.cards), "cards")
	for _,v := range p1.cards {
		fmt.Println(v)
	}
	fmt.Println(len(p2.cards), "cards")
	fmt.Println("========")
	for _,v := range p2.cards {
		fmt.Println(v)
	}

}