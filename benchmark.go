package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"strings"
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
	// prime_time := int64(5000000000)
	sample_time := int64(60000000000)
	// sample_time := int64(10000000000)
	// sample_time := int64(5000000000)
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
				phase = 3
			}

		} else if phase == 4 {
			break monitor
		}

		if phase == 2 && (current_time - last_sample_time) > sample_frequency {
			last_sample_time = current_time
			samples = append(samples, speed)
		}

		if (current_time - last_display_time) > display_frequency {
			last_display_time = current_time

			if phase == 1 {
				fmt.Printf("\r%d. priming | et = %ds; g = %d; s = %.5f g/ms; \t",
				phase, elapsed_time / ns, total_games, speed_v)
			} else if phase == 2 {
				fmt.Printf("\r%d. sampling | et = %ds; g = %d; s = %.5f g/ms; t = %d; \t",
				phase, elapsed_time / ns, total_games, speed_v, len(samples))				
			} else if phase == 3 {
				phase = 4
				fmt.Printf("\r%d. done | et = %ds; g = %d; s = %.5f g/ms; t = %d; \t",
				phase, elapsed_time / ns, total_games, speed_v, len(samples))	
			}
			
		}

	}

	// final statistics
	mean = get_mean(samples)
	stdev = get_standard_deviation(samples, mean)
	cov = get_coefficient_of_variation(mean, stdev)

	min_max_delta := maximum_speed - minimum_speed

	one_sigma_lower := (mean-stdev)*float64(ms)
	one_sigma_upper := (mean+stdev)*float64(ms)
	one_sigma_delta := one_sigma_upper - one_sigma_lower

	const t_score = 3.291 // 99.9%
	const one_percent = .01

	nfci_lower := (mean - (t_score * (stdev / math.Sqrt(float64(len(samples)))))) * float64(ms)
	nfci_upper := (mean + (t_score * (stdev / math.Sqrt(float64(len(samples)))))) * float64(ms)
	nfci_delta := nfci_upper - nfci_lower

	points := make([]string, 0, 3)

	if cov < one_percent {
		points = append(points, "1%cov")
	}

	if one_sigma_lower < speed_v && speed_v < one_sigma_upper {
		points = append(points, "1σ")
	}

	if nfci_lower < speed_v && speed_v < nfci_upper {
		points = append(points, "99.9%CI")
	}

	
	fmt.Println("\n---\n")

	fmt.Printf("Samples: %d collected\n", len(samples))
	fmt.Printf("Mean: %.5f\n", mean * float64(ms))
	fmt.Printf("Standard Deviation: %.5f\n", stdev * float64(ms))
	fmt.Printf("Coefficient of Variation: %.5f\n", cov)

	fmt.Printf("Min-Max:\t < %.5f - %.5f > Δ %.5f\n", minimum_speed*float64(ms), maximum_speed*float64(ms), min_max_delta*float64(ms))
	
	fmt.Printf("1-σ:\t\t < %.5f - %.5f > Δ %.5f\n", one_sigma_lower, one_sigma_upper, one_sigma_delta)

	fmt.Printf("99.9%% CI:\t < %.5f - %.5f > Δ %.5f", nfci_lower, nfci_upper, nfci_delta)


	fmt.Println("\n---\n")
	
	fmt.Printf("Threads: %d\n", threads)
	fmt.Printf("Speed: %.5f\n", speed_v)
	fmt.Printf("Total Games: %d\n", total_games)
	fmt.Printf("Elapsed Time: %.0f seconds\n", float64(elapsed_time / ns))

	fmt.Printf("Rank Passes: %s\n", rank_reason(points))
	fmt.Printf("\nScore: %d %s\n", math_round(speed_v), rank(points))

}

func rank(passes []string) string {
	v := len(passes)
	letter := ""
	switch v {
		case 3: letter = "A"
		case 2: letter = "B"
		case 1: letter = "C"
		default: letter = "D"
	}
	return letter
}

func rank_reason(passes []string) string {
	reason := ""
	if len(passes) == 0 {
		reason = "none"
	} else {
		reason = strings.Join(passes[:], ", ")
	}
	return reason
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
				// no data available
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
	return stdev / mean
}
