package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"strings"
	"sort"
)

const ms int64 = 1000000
const ns int64 = 1000000000

// Benchmark accepts a number of threads,
// and will eventually benchmark.
func Benchmark(threads int, multiplier float64) {

	progress_channels := make([](chan int), threads)

	create_threads(threads, &progress_channels)

	// 1/15 of a second
	const display_frequency int64 = ns/10
	// 1/200 of a second
	const sample_frequency int64 = 	ns/200

	// 10 seconds
	var prime_time int64 = 	10000000000
	// 50 seconds
	var sample_time int64 = 50000000000

	// in this way, the total benchmark time is 60 seconds
	// and any integer multipliers will be multiples of 60

	if multiplier != 1.0 {
		prime_time = int64(float64(prime_time) * multiplier)
		sample_time = int64(float64(sample_time) * multiplier)
	}

	// development times
	// const prime_time int64 = 	5000000000
	// const sample_time int64 = 5000000000

	// when to end the the benchmark
	var end_time int64 = prime_time + sample_time

	var sample_size int64 = sample_time / sample_frequency

	var samples []float64 = make([]float64, 0, sample_size)

	var start_time int64 = time.Now().UnixNano()
	var current_time int64 = time.Now().UnixNano()
	var elapsed_time int64

	var last_display_time int64 = current_time
	var last_sample_time int64 = current_time

	var phase int = 1

	// to avoid the rare but possible case when `collect_progress` has not
	// counted any games finishing yet
	var total_games int64 = 1

	var speed float64 = 0

	var maximum_speed float64
	var minimum_speed float64

	monitor: for true {
		total_games += int64(collect_progress(&progress_channels))

		current_time = time.Now().UnixNano()
		elapsed_time = current_time - start_time

		speed = float64(total_games) / float64(elapsed_time)

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
				phase, elapsed_time / ns, total_games, speed * float64(ms))
			} else if phase == 2 {
				fmt.Printf("\r%d. sampling | et = %ds; g = %d; s = %.5f g/ms; t = %d; \t",
				phase, elapsed_time / ns, total_games, speed * float64(ms), len(samples))
			} else if phase == 3 {
				phase = 4
				// intentionally blank line
				fmt.Printf("\r%d. done                                                                 \t",
				phase)
			}
		}
	}

	// constants
	const t_score = 3.291 // 99.9% t-score
	const one_percent = .01 // 1%
	const ten_percent = .1 // 10%

	// final statistics
	var mean float64 = get_mean(samples)
	var median float64 = get_median(samples)
	var stdev float64 = get_standard_deviation(samples, mean)
	var cov float64 = get_coefficient_of_variation(mean, stdev)

	// the signed value delta of mean and median
	var mean_median_delta float64 = math.Abs(median - mean)
	var mm_lower, mm_upper float64 = math.Min(mean, median), math.Max(mean, median)

	// the delta of the max-min speeds
	var min_max_delta float64 = maximum_speed - minimum_speed
	var max_ten_percent float64 = maximum_speed * ten_percent

	// one_sigma is 1 standard deviation away from the mean
	var one_sigma_lower float64 = (mean-stdev)
	var one_sigma_upper float64 = (mean+stdev)
	var one_sigma_delta float64 = one_sigma_upper - one_sigma_lower

	// 99.9% confidence interval; how likely it is that the true mean lies within
	var ci_lower float64 = (mean - (t_score * (stdev / math.Sqrt(float64(len(samples))))))
	var ci_upper float64 = (mean + (t_score * (stdev / math.Sqrt(float64(len(samples))))))
	var ci_delta float64 = ci_upper - ci_lower

	// controversial section! points are given
	// based on passing basic statistical testing criteria
	var criteria map[string]bool = make(map[string]bool)

	// pass: the mean_median_delta is less than the standard deviation
	criteria["1"] = mean_median_delta < stdev

	// pass: is the delta smaller than 10% of the max?
	criteria["2"] = min_max_delta < max_ten_percent

	// pass: COV < 1%; stdev / mean
	criteria["3"] = cov < one_percent

	// pass: the final speed is within 1 stdev
	criteria["4"] = one_sigma_lower < speed && speed < one_sigma_upper

	// pass: the final speed is near the true mean; within the confidence interval
	criteria["5"] = ci_lower < speed && speed < ci_upper

	// only printing below
	// 1. raw statisitics
	// 2. ranges
	// 3. summary of the benchmark
	// 4. rank
	// 5. score

	fmt.Printf("\n---\n")

	fmt.Printf("Samples: %9d\n", len(samples))
	fmt.Printf("Mean:\t %9.5f\n", toms(mean))
	fmt.Printf("Median:\t %9.5f\n", toms(median))
	fmt.Printf("S.D.:\t %9.5f\n", toms(stdev))
	fmt.Printf("C.O.V.:\t %9.5f\n", cov)

	fmt.Printf("---\n")

	fmt.Printf("Min-Max:\t < %9.5f - %9.5f > Δ %9.5f\n",
		toms(minimum_speed),
		toms(maximum_speed),
		toms(min_max_delta))

	fmt.Printf("1-σ:\t\t < %9.5f - %9.5f > Δ %9.5f\n",
		toms(one_sigma_lower),
		toms(one_sigma_upper),
		toms(one_sigma_delta))

	fmt.Printf("μ-Median:\t < %9.5f - %9.5f > Δ %9.5f\n",
		toms(mm_lower),
		toms(mm_upper),
		toms(mean_median_delta))

	fmt.Printf("99.9%% CI:\t < %9.5f - %9.5f > Δ %9.5f\n",
		toms(ci_lower),
		toms(ci_upper),
		toms(ci_delta))


	fmt.Printf("---\n")

	fmt.Printf("Threads: %d\n", threads)
	fmt.Printf("Multiplier: %.2f\n", multiplier)
	fmt.Printf("Speed: %.5f g/ms\n", toms(speed))
	fmt.Printf("Games: %d\n", total_games)
	fmt.Printf("Duration: %.1fs\n", float64(elapsed_time / ns))

	fmt.Printf("---\n")

	fmt.Printf("Rank: (%d/%d) %s\n", rank_passes(criteria), len(criteria), rank_letter(criteria))
	fmt.Printf("Rank Criteria: %s\n", rank_reason(criteria))

	fmt.Printf("---\n")

	fmt.Printf("Score: %d\n", math_round(toms(speed)))

}

// Instead of multiplying all over, wrap float's that need to be
// represented in milliseconds with this function.
func toms(f float64) float64 {
	return f * float64(ms)
}

// Rank letter accepts a list of passed tests that
// define certain statistical qualities.
// For each successful pass, a better letter rank is returned.
func rank_letter(criteria map[string]bool) string {
	v := rank_passes(criteria)
	letter := ""
	switch v {
		case 5: letter = "A+"
		case 4: letter = "A"
		case 3: letter = "B"
		case 2: letter = "C"
		case 1: letter = "D"
		default: letter = "F"
	}
	return letter
}

// Rank passes totals the number of true keys in the given criteria array
// and returns that as an integer.
func rank_passes(criteria map[string]bool) int {
	r := 0
	for _, v := range criteria {
		if v {
			r++
		}
	}
	return r;
}

// Rank reason concatenates a string, or reports none.
// The rank reasons are set in a string slice.
func rank_reason(criteria map[string]bool) string {
	reason := ""
	passes := rank_passes(criteria)
	if passes == 0 {
		reason = "none"
	} else {
		passed := []string{}
		for k,v := range criteria {
			if v {
				passed = append(passed, k)
			}
		}
		reason = strings.Join(passed[:], " | ")
	}
	return reason
}

// Math round attempts to round a float to an integer.
func math_round(f float64) int64 {
	return int64(math.Floor(f + .5))
}

// Create threads accepts a number of threads and a channel pointer,
// which will be populated with channels for each respective thread
// spun up. Each thread runs a game loop, and upon completion of each game,
// the thread sends through the progress channel.
func create_threads(threads int, channels *[](chan int)) {
	for i := 0; i < threads; i++ {

		progress := make(chan int, 4096)
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

// Collect progress accesses each channel and listens, in a non-blocking fashion,
// for any data passed back to it that indicates progress has been made.
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

func get_median(samples []float64) float64 {
	sort.Float64s(samples)
	var length int = len(samples)
	var median float64 = 0
	if length % 2 == 0 {
		a := samples[length / 2 - 1]
		b := samples[length / 2 + 1]
		median = (a+b)/2
	} else {
		median = samples[length / 2]
	}
	return median
}

// Get mean calculates the mean based on the given samples.
func get_mean(samples []float64) float64 {
	var total float64 = 0
	for _, v := range samples {
		total += v
	}
	var mean float64 = total / float64(len(samples))
	return mean
}

// Get standard deviation calculates the stdev based on the given samples and the mean.
// TODO: implement `online_variance` algorithm
func get_standard_deviation(samples []float64, mean float64) float64 {
	var total float64 = 0
	for _, v := range samples {
		total += math.Pow(v - mean, 2)
	}
	var stdev float64 = math.Sqrt(total / float64(len(samples) - 1))
	return stdev
}

// Get coefficient of variation calculations the standard deviation to mean ratio.
func get_coefficient_of_variation(mean float64, stdev float64) float64 {
	return stdev / mean
}
