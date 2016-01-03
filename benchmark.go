package main

import (
	"fmt"
	"time"
	"math/rand"
)

func Benchmark() {

	threads := 4

	progress := make(chan int, 4096)

	for i := 0; i < threads; i++ {

		go func() {
			source := rand.NewSource(time.Now().UnixNano())
			generator := rand.New(source)
			for true {
				Game(generator)
				progress <- 1
			}
		}()

	}

	// samples := make([]float64, 10000)

	phase := 1
	_ = phase

	total_games := int64(1)

	start_time := time.Now().UnixNano()
	current_time := time.Now().UnixNano()

	prime_time := int64(60000000000)
	// maximum_tests := 240

	display_frequency := int64(50000000)
	// sample_frequency := 5000000

	last_display_time := int64(0)
	// last_sample_time := 0

	ms := int64(1000000)
	ns := int64(1000000000)

	elapsed_time := int64(0)
	test_time := int64(0)

	test_started := false

	speed := 0.0
	speed_v := 0.0
	_ = speed_v
	rate := 0.0

	maximum_speed := float64(0)
	// mean := 0.0
	// stdev := 0.0
	// cov := 0.0

	monitor: for true {
		select {
			case p := <-progress:
				total_games += int64(p)
			default:
				// do nothing!
		}

		current_time = time.Now().UnixNano()
		elapsed_time = current_time - start_time

		rate = float64(elapsed_time) / float64(total_games)

		speed = 1.0 / rate
		speed_v = speed * float64(ms)

		if maximum_speed < speed {
			maximum_speed = speed
		}

		if !test_started && elapsed_time >= prime_time {
			// phase 1
			test_started = true
			phase = 2
		} else if test_started && elapsed_time >= test_time {
			
			break monitor

		}

		if (current_time - last_display_time) > display_frequency {
			last_display_time = current_time
			
			fmt.Printf("\r%d. et = %ds; g = %d; s = %f g/ms; \t",
				phase, elapsed_time / ns, total_games, speed_v)
			
		}

	}

	fmt.Println("count:", total_games)


}

