package main

import (
	"testing"
)

func BenchmarkGame(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Game()
	}
}