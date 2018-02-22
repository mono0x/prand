package prand

import (
	"math/rand"
	"testing"
	"time"
)

func TestSource(t *testing.T) {
	source := NewSource(DefaultNew)
	_ = source.Int63()
}

func TestInt63(t *testing.T) {
	_ = Int63()
}

func TestUint64(t *testing.T) {
	_ = Uint64()
}

func BenchmarkMathRandInt63(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rand.Int63()
	}
}

func BenchmarkPrandInt63(b *testing.B) {
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Int63()
	}
}

func BenchmarkMathRandInt63Parallel(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	b.SetBytes(8)
	b.SetParallelism(64)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = rand.Int63()
		}
	})
}

func BenchmarkPrandInt63Parallel(b *testing.B) {
	b.SetBytes(8)
	b.SetParallelism(64)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Int63()
		}
	})
}

func BenchmarkMathRandUint64(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rand.Uint64()
	}
}

func BenchmarkPrandUint64(b *testing.B) {
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Uint64()
	}
}

func BenchmarkMathRandUint64Parallel(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	b.SetBytes(8)
	b.SetParallelism(64)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = rand.Uint64()
		}
	})
}

func BenchmarkPrandUint64Parallel(b *testing.B) {
	b.SetBytes(8)
	b.SetParallelism(64)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Uint64()
		}
	})
}
