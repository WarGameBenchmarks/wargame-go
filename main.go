package main

import (
	"fmt"
	"runtime"
)

const DEVELOPMENT = false

func main() {

	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)
	
	if DEVELOPMENT == true {
		fmt.Println("WarGame Go")
		fmt.Println("\tSee `debug.log` for details.")
	} else {
		fmt.Println("WarGame Go")
	}

	Benchmark()

}