package main

import (
	"fmt"
	"runtime"
)

func main() {

	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)
	
	fmt.Println("WarGame Go")
	
	Benchmark()

}