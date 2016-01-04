package main

import (
	"os"
	"strconv"
	"fmt"
	"runtime"
)

func main() {

	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	arguments := os.Args[1:]
	
	threads := 1

	if len(arguments) >= 1 {
		fmt.Println(arguments[0])
		i, err := strconv.Atoi(arguments[0])
		if err != nil {
			fmt.Println(err)
		} 
		threads = i
	}

	fmt.Println("WarGame Go")

	fmt.Printf("settings: threads = %d\n\n", threads)
	
	Benchmark(threads)

}