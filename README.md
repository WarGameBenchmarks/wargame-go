WarGame Go
==========

This is a [Go](https://golang.org/)lang port of [@ryanmr](http://twitter.com/ryanmr)'s infamous WarGame *kinda-sorta* benchmark.

**Disclaimer:** This is a *kinda-sorta* benchmark and as such, it should not be taken too seriously.

Changelog
---------

You can view the [changelog](changelog.md), that hopefully mentions major differences between versions.

Legend
------

You will definitely want to read WarGame's output &mdash; there are descriptions and discussions of the methods and statistics used are available in the [legend](https://github.com/WarGameBenchmarks/wargame/blob/master/legend.md).

How To Run
----------

To run the executable directly:

```
./wargame-go [number of threads] [multiplier]
```

- *number of threads*: how many *go routines* or threads should the benchmark use
- *multiplier*: how much less or additional time should the default (60 seconds) duration be altered; this is multiplicative.
  - 1x: 10s priming, 50s sampling
  - 1.5x: 15s priming, 75s sampling
  - 3x: 30s priming, 90s sampling

If the arguments are not specified, they will both default to 1.

How To Compile
--------------

Clone this repository to your *Go workspace* and install it.

```
mkdir -p ~/{go workspace}/src/github.com/ryanmr/wargame-go/
go install
```

Alternatively, run this inside of the repository's directory:

```
go build
```

You can use the generated binary:

```
./wargame-go
```

Sample Output
-------------

```
➜  wargame-go git:(release-alpha) go build        
➜  wargame-go git:(release-alpha) ./wargame-go 4        
WarGame Go
settings: threads = 4; multiplier = 1.00

4. done                                                                 	
---
Samples:      8651
Mean:	  13.38033
Median:	  13.49081
S.D.:	   0.23374
C.O.V.:	   0.01747
---
μ-Median:	 <  13.38033 -  13.49081 > Δ   0.11048
Min-Max:	 <  12.67879 -  13.69475 > Δ   1.01596
1-σ:		 <  13.14659 -  13.61407 > Δ   0.46747
99.9% CI:	 <  13.37206 -  13.38860 > Δ   0.01654
---
Threads: 4
Multiplier: 1.00
Speed: 13.44045 g/ms
Games: 807112
Duration: 60.0s
---
Rank: (3/5) B
Rank Criteria: 4 | 1 | 2
---
Score: 13
```
