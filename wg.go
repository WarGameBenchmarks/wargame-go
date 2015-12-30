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

func main() {
	vs := []Value{
		Two, Three, Four, 
		Five, Six, Seven, 
		Eight, Nine, Ten, 
		Jack, Queen, King, Ace}
	for _,v := range vs {
		fmt.Println(v)
	}
}