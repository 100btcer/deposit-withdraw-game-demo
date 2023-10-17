package rand_test

import (
	mathrand "math/rand"
	"music-nft/internal/utils/rand"
	"testing"
)

var (
	buf32   = make([]byte, 32)
	buf1024 = make([]byte, 1024)
)

func Benchmark_RandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.RandString(32)
	}
}

func Benchmark_RandStringEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.RandStringEx(32)
	}
}

func Benchmark_MathrandRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathrand.Read(buf32)
	}
}

func Benchmark_CRandRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.B(32)
	}
}

func Benchmark_CryRand_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathrand.Int()
	}
}
