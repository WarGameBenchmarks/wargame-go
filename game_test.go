package main

import (
	"testing"
	"time"
	"math/rand"
)

func BenchmarkGame(b *testing.B) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	for n := 0; n < b.N; n++ {
		Game(generator)
	}
}