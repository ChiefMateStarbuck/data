package main

import (
	"testing"
)

func benchmarkObject(b *testing.B, size int) {
	ants := AoS(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountAoS(ants)
	}
}

func benchmarkData(b *testing.B, size int) {
	ants := SoA(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountSoA(ants)
	}
}

func Benchmark100Object(b *testing.B) { benchmarkObject(b, 100) }
func Benchmark100Data(b *testing.B)   { benchmarkData(b, 100) }

func Benchmark1000Object(b *testing.B) { benchmarkObject(b, 1000) }
func Benchmark1000Data(b *testing.B)   { benchmarkData(b, 1000) }

func Benchmark10000Object(b *testing.B) { benchmarkObject(b, 10000) }
func Benchmark10000Data(b *testing.B)   { benchmarkData(b, 10000) }

func Benchmark100000Object(b *testing.B) { benchmarkObject(b, 100000) }
func Benchmark100000Data(b *testing.B)   { benchmarkData(b, 100000) }
