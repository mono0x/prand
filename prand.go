package prand

import (
	"bytes"
	cryptorand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"sync"
	"unsafe"
)

type source struct {
	pool sync.Pool
}

var _ rand.Source64 = (*source)(nil)

func NewSource(newSource func() rand.Source) rand.Source {
	return &source{
		pool: sync.Pool{
			New: func() interface{} { return newSource() },
		},
	}
}

func (s *source) Seed(seed int64) {}

func (s *source) Int63() int64 {
	source := s.pool.Get().(rand.Source)
	defer s.pool.Put(source)
	return source.Int63()
}

func (s *source) Uint64() uint64 {
	source := s.pool.Get().(rand.Source)
	defer s.pool.Put(source)
	if s64, ok := source.(rand.Source64); ok {
		return s64.Uint64()
	}
	return uint64(source.Int63())>>31 | uint64(source.Int63())<<32
}

func DefaultNew() rand.Source {
	b := make([]byte, unsafe.Sizeof(int64(0)))
	if _, err := cryptorand.Read(b); err != nil {
		panic(err) // TODO: improve error handling
	}
	var seed int64
	if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &seed); err != nil {
		panic(err) // TODO: improve error handling
	}
	return rand.NewSource(seed).(rand.Source)
}

var globalRand = rand.New(NewSource(DefaultNew))

func Int63() int64 { return globalRand.Int63() }

func Uint32() uint32 { return globalRand.Uint32() }

func Uint64() uint64 { return globalRand.Uint64() }

func Int31() int32 { return globalRand.Int31() }

func Int() int { return globalRand.Int() }

func Int63n(n int64) int64 { return globalRand.Int63n(n) }

func Int31n(n int32) int32 { return globalRand.Int31n(n) }

func Intn(n int) int { return globalRand.Intn(n) }

func Float64() float64 { return globalRand.Float64() }

func Float32() float32 { return globalRand.Float32() }

func Perm(n int) []int { return globalRand.Perm(n) }

func Shuffle(n int, swap func(i, j int)) { globalRand.Shuffle(n, swap) }

func Read(p []byte) (n int, err error) { return globalRand.Read(p) }

func NormFloat64() float64 { return globalRand.NormFloat64() }

func ExpFloat64() float64 { return globalRand.ExpFloat64() }
