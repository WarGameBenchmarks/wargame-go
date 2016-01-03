package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func Benchmark() {

	threads := 8

	progress_channels := make([](chan int), threads)

	create_threads(threads, &progress_channels)

	samples := make([]float64, 0, 12000)

	phase := 1

	total_games := int64(1)

	start_time := time.Now().UnixNano()
	current_time := time.Now().UnixNano()

	prime_time := int64(10000000000)
	// sample_time := int64(60000000000)
	sample_time := int64(10000000000)
	end_time := prime_time + sample_time

	display_frequency := int64(50000000)
	sample_frequency := int64(5000000)

	last_display_time := int64(0)
	last_sample_time := int64(0)

	ms := int64(1000000)
	ns := int64(1000000000)

	elapsed_time := int64(0)

	speed := 0.0
	speed_v := 0.0
	rate := 0.0

	var maximum_speed float64
	var minimum_speed float64
	mean := 0.0
	stdev := 0.0
	cov := 0.0

	monitor: for true {		
		total_games += int64(collect_progress(&progress_channels))

		current_time = time.Now().UnixNano()
		elapsed_time = current_time - start_time

		rate = float64(elapsed_time) / float64(total_games)

		speed = 1.0 / rate
		speed_v = speed * float64(ms)


		if phase == 1 && elapsed_time >= prime_time {
			phase = 2

			maximum_speed = speed
			minimum_speed = speed

		} else if phase == 2 {

			if maximum_speed < speed {
				maximum_speed = speed
			}

			if minimum_speed > speed {
				minimum_speed = speed
			}

			if elapsed_time >= end_time {
				break monitor
			}

		}

		if phase == 2 && (current_time - last_sample_time) > sample_frequency {
			last_sample_time = current_time
			samples = append(samples, speed)
		}

		if (current_time - last_display_time) > display_frequency {
			last_display_time = current_time

			mean = get_mean(samples)
			stdev = get_standard_deviation(samples, mean)
			cov = get_coefficient_of_variation(mean, stdev)
			
			if phase == 1 {
				fmt.Printf("\r%d. priming | et = %ds; g = %d; s = %.5f g/ms; \t",
				phase, elapsed_time / ns, total_games, speed_v)
			} else if phase == 2 {
				fmt.Printf("\r%d. sampling | et = %ds; g = %d; s = %.5f g/ms; t = %d; \t",
				phase, elapsed_time / ns, total_games, speed_v, len(samples))				
			}
			
		}

	}


	// final statistics
	mean = get_mean(samples)
	stdev = get_standard_deviation(samples, mean)
	cov = get_coefficient_of_variation(mean, stdev)
	
	fmt.Println("\n---\n")

	fmt.Printf("Samples: %d collected\n", len(samples))
	fmt.Printf("Mean: %.5f\n", mean * float64(ms))
	fmt.Printf("Standard Deviation: %.5f\n", stdev * float64(ms))
	fmt.Printf("Coefficient of Variation: %.5f\n", cov)

	fmt.Printf("Minimum Speed: %.5f g/ms\n", minimum_speed * float64(ms))
	fmt.Printf("Maximum Speed: %.5f g/ms\n", maximum_speed * float64(ms))
	
	fmt.Printf("Distribution:\t < %.2f | %.2f | %.2f >\n", (mean-stdev)*float64(ms), speed_v, (mean+stdev)*float64(ms))
	fmt.Printf("Min-Max:\t < %.2f | %.2f | %.2f >\n", math.Abs(minimum_speed-speed)*float64(ms), speed_v, math.Abs(maximum_speed-speed)*float64(ms))

	fmt.Printf("95%% CI:\t < %.2f | %.2f | %.2f >", 
		(mean - 1.960 * (stdev / math.Sqrt(float64(len(samples))))) * float64(ms),
		speed_v,
		(mean + 1.960 * (stdev / math.Sqrt(float64(len(samples))))) * float64(ms))


	fmt.Println("\n---\n")
	
	fmt.Printf("Threads: %d\n", threads)
	fmt.Printf("Speed: %.5f\n", speed_v)
	fmt.Printf("Total Games: %d\n", total_games)
	fmt.Printf("Elapsed Time: %d nanoseconds; %.0f seconds\n",
		elapsed_time, float64(elapsed_time / ns))

	fmt.Printf("\nScore: %d\n", math_round(speed_v))

}

func math_round(f float64) int64 {
	return int64(math.Floor(f + .5))
}

func create_threads(threads int, channels *[](chan int)) {
	for i := 0; i < threads; i++ {

		progress := make(chan int, threads * 1024)
		(*channels)[i] = progress

		go func() {
			source := rand.NewSource(time.Now().UnixNano())
			generator := rand.New(source)
			for true {
				Game(generator)
				progress <- 1
			}
		}()

	}
}

func collect_progress(channels *[](chan int)) int {
	r := 0	
	for _,v := range *channels {
		select {
			case p := <-v:
				r += p
			default:
				// do nothing!
		}
	}
	return r
}

func get_mean(samples []float64) float64 {
	var total float64 = 0
	for _, v := range samples {
		total += v
	}
	var mean float64 = total / float64(len(samples))
	return mean
}


// TODO: implement `online_variance` algorithm
func get_standard_deviation(samples []float64, mean float64) float64 {
	var total float64 = 0
	for _, v := range samples {
		total += math.Pow(v - mean, 2)
	}
	var stdev float64 = math.Sqrt(total / float64(len(samples)))
	return stdev
}

func get_coefficient_of_variation(mean float64, stdev float64) float64 {
	return (stdev / mean) * float64(100)
}
