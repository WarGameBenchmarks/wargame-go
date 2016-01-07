Legend
======

The WarGame plays endless loops of the War card game in the number of specified threads.

The legend contains descriptions you may see in the WarGame output.

#### Threads

The number of threads that parallel games of War will be run in.

Often, threads are implemented differently across operating systems and language platforms.

#### Multiplier

How much to multiply the base timings.

Examples:

- 1x: 10s priming, 50s sampling
- 1.5x: 15s priming, 75s sampling
- 3x: 30s priming, 90s sampling
- 5x: 50s priming, 250s sampling

Using `./wargame-go 4 5` will use 4 threads and the benchmark will run for 5 minutes, of which 50 seconds will be prime time, and 250 seconds will be sample time.

#### Prime Time (priming)

To get the benchmark warmed up, it is primed for this duration before statistical sampling begins.

The benchmark records various initial times with a high precision timer. To amortize the initialization costs of variables, loops, and other operations, the *prime time* spreads out the cost over a known period of time before sampling begins to decrease the chance that results are skewed.

By default, *prime time* is 10 seconds. This value can be increased with the *multiplier*.



#### Sample Time (sampling)

Statistics are collected every 5 milliseconds during the *sample time*. The value being collected is the *speed* at which the benchmark is running games.

By default, *sample time* is 50 seconds. This value can be increased with the *multiplier*.

#### Elapsed Time (et)

This is the time that has passed since the benchmark began.

#### Games (g)

The number of war games counted so far.

#### Speed (s)

The speed measured so far, defined in *games per millisecond* units.

#### Samples

The number of samples collected during the sampling time specified. Each sample is a instantaneous measurement of the speed of the benchmark in *games per millisecond* units.

#### Mean

The mean is the arithmetic average of the samples collected.

#### Median

The median is the *middle* most sample collected.

#### S.D.

Standard Deviation is the average distance that the samples are a part.

#### C.O.V.

[Coefficient of Variation](https://en.wikipedia.org/wiki/Coefficient_of_variation) is a unitless ratio that measures `standard deviation / mean`.

#### μ-Median

This is the range between the mean and the median.

#### μ-Median Δ

This is the *delta* of the mean and median.

#### Min-Max

This is the range between the smallest and largest `speed` values that appeared while running in the *sample time*, though not limited only to the samples collected.

#### Min-Max Δ

This is the *delta* of the minimum and maximum speeds recorded during the *sample time*.

#### 1-σ

This is the range defined by one standard deviation (or one-sigma) away from the mean. [Typically](https://en.wikipedia.org/wiki/68%E2%80%9395%E2%80%9399.7_rule), 1 standard deviation in a normal distribution contains 68%.

#### 1-σ Δ

This is the *delta* of the bounds formed by `mean - standard deviation` and `mean - standard deviation`.

#### 99.9% CI

The [confidence interval](https://en.wikipedia.org/wiki/Confidence_interval) represents a range that should contain the true mean parameter 99.9% of the time if the experiment were repeated.

#### 99.9% CI Δ

This is the *delta* of the bounds formed [confidence interval](https://en.wikipedia.org/wiki/Student%27s_t-distribution#Confidence_intervals). The formula is based on the following:

```
mean +/- 3.021 * (standard deviation / sqrt(sample size))
```

In the formula above, `3.021` represents the *t-score* value from the t-table, using a 99.9 probability and *"infinite"* degrees of freedom.

#### Rank

*Rank* is an analysis of some of the samples using rudimentary statistics.

Disclaimer: [@ryanmr](http://twitter.com/ryanmr) is not a statistician. He is 95% sure about this fact.

For each *rank criteria* met, the benchmark earns a progressively better grade letter: F, D, C, B, A, and A+.

- 0/5 = F
- 1/5 = D
- 2/5 = C
- 3/5 = B
- 4/5 = A
- 5/5 = A+

The *rank criteria* below are labeled numerically, and they are likewise label as such in the output.

In testing, generally though not strictly, the lower criteria are easier to meet, while the larger criteria are harder to meet.

The *rank* is entirely partially subjective by design. Lacking rigorous statistical analysis, it is an easy way to determine if a sample set is poor, fair, good or great, but it nothing more than that.

The rank criteria listed below are *arbitrary*, and as such, they should be considered very carefully before they are taken seriously.

#### Rank Criteria

The rank criteria were designed to be progressively more difficult to achieve, but need not be achievable in a specific order.

The *Rank Criteria* output may be shown like so:

```
Rank Criteria: 4 | 1 | 2 | 3 | 5
```

This means that these criteria were met. The order *does not* matter.

```
Rank Criteria: none
```

This means that no criteria were met.

The follow is a discussion on each criteria. The numbers shown in the output correspond to the numbers described below.

##### 1 `|mean-median| < standard deviation`

The mean is the average of all samples, and as such, outliers can pull the mean towards an extreme. The median is less prone to outliers bias. The *delta* of the mean and median is likely to be small for most sample sets, and if that *delta* is smaller than the *standard deviation*, the mean-median difference should not be significant.

##### 2 `|min-max| < (10% max)`

The minimum and maximum are bounded by the sample time's smallest and largest speeds. The *delta* of the minimum and maximum shows the range that all sample values could be.

Arbitrarily, 10% of the maximum is rule of thumb for the *delta* to be under. This helps to restrict the ranges that may produce desirable results.

This is an example of this criteria being met:

```
Min-Max:	 <  12.67879 -  13.69475 > Δ   1.01596
```

In this case, `1.01596` is less than `1.369475`, so the criteria is met.


This is an example of this criteria failing:

```
Min-Max:	 <   6.25831 -  15.30748 > Δ   9.04917
```

This range is huge. `9.04917` clearly exceeds the 10% maximum threshold of `1.530748`, and thus this sample set is not awarded.

##### 3 `standard deviation < (%1 mean)`

Arbitrarily, this criteria asks the *standard deviation* to be 1% of the the *mean*. While the *coefficient of variation* is a unitless value, it easily shows when the standard deviation is fairly small compared to its accompanying mean.

This example meets the criteria:

```
Mean:	  23.78822
S.D.:	   0.03400
```

In this case, `0.03400 / 23.78822` is `0.00142` which is much smaller than `.01`.

In this example, the C.O.V. does not show that the *standard deviation* is 1% of the *mean*, and thus does not meet the criteria.

```
Mean:	  10.77338
S.D.:	   3.36117
```

`3.36117/10.77338` is `0.31198` which exceeds `.01`, and thus fails.

##### 4 `μ-σ < speed < μ-σ`

This measures the final speed against the bounds of one standard below and above the mean. Typically, as mentioned, this relates to 68% of the sample data over a normal distribution. It should be fairly common for the final speed to fall into this range.

This is an example where the final speed meets the criteria:

```
1-σ:		 <  13.14659 -  13.61407 > Δ   0.46747
Speed: 13.44045 g/ms
```

This is an example of where the final speed does not meet the criteria:

```
1-σ:		 <   7.41222 -  14.13455 > Δ   6.72233
Speed: 6.25703 g/ms
```

In the first case, the range is small and the final speed is within the range. In the latter, the range does not include the final speed.

##### 5 `99.9%` Confidence Interval

The formula for the confidence interval is:

```
mean +/- 3.021 * (standard deviation / sqrt(sample size))
```

This criteria asks the final speed to be near the range of the true mean. Recall that the confidence interval does not say much about the parameter mean, but instead is used to compare against other similar sample sets. 99.9% of sample sets should have a true mean described by the ranges yielded across many experiments. There is no reason to ask the final speed to be in this range. In this way, this criteria is *arbitrary*.

But if the final speed is in this range, it is also likely near the true mean. This attribute is quite rare in testing. Effectively, this is *extra credit*.

This is an incredibly rare run:

```
99.9% CI:	 < 14.19066 - 14.20496 > Δ 0.01430
Speed: 14.19112 g/ms
```

In this run, the final speed is narrowly inside the 99.9% confidence interval range.

```
99.9% CI:	 <  10.61822 -  10.92855 > Δ   0.31033
Speed: 6.25703 g/ms
```

In this run, the final speed is no where near the range, let alone inside of it.
