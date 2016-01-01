package main

import (
	"fmt"
)

const DEVELOPMENT = false

func main() {

	
	if DEVELOPMENT == true {
		fmt.Println("WarGame Go")
		fmt.Println("\tSee `debug.log` for details.")
	} else {
		fmt.Println("WarGame Go")
	}

	Game()

}