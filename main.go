package main

import (
	"os"
	"strconv"
	"fmt"
	"math"
	"runtime"
)

func main() {

	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	arguments := os.Args[1:]

	threads := 1
	multiplier := 1.0

	if len(arguments) == 2 {
		i1, err1 := strconv.Atoi(arguments[0])
		i2, err2 := strconv.ParseFloat(arguments[1], 64)
		switch {
			case err1 != nil:
				fmt.Println(err1)
			case err2 != nil:
				fmt.Println(err2)
			default:
				threads = i1
				multiplier = math.Abs(i2)
		}
	} else if len(arguments) == 1 {
		i, err := strconv.Atoi(arguments[0])
		if err != nil {
			fmt.Println(err)
		} else {
			threads = i
		}
	}

	fmt.Println("WarGame Go")

	fmt.Printf("settings: threads = %d; multiplier = %.2f\n\n", threads, multiplier)

	Benchmark(threads, multiplier)

}
