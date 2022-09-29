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

var (
	_ rand.Source64 = (*source)(nil)
	_ Source        = (*newSource)(nil)
	_ Default       = (*defaultSource)(nil)
)

type newSource struct {
	Source func() rand.Source
}

type defaultSource struct {
	Error error
}

type Source interface {
	New() rand.Source
}

type Default interface {
	New() rand.Source
}

func (new newSource) New() rand.Source {
	return &source{
		pool: sync.Pool{
			New: func() interface{} {
				return new.Source()
			},
		},
	}
}

func (s *defaultSource) New() rand.Source {
	b := make([]byte, unsafe.Sizeof(int64(0)))
	if _, err := cryptorand.Read(b); err != nil {
		s.Error = err
		return nil
	}
	var seed int64
	if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &seed); err != nil {
		s.Error = err
		return nil
	}
	return rand.NewSource(seed)
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

var (
	defaultsrc = defaultSource{
		Error: nil,
	}
	src = newSource{
		Source: defaultsrc.New,
	}
	globalRand = rand.New(src.New())
	err        = defaultsrc.Error
)

func Int63() (int64, error) {
	if err != nil {
		return globalRand.Int63(), err
	}
	return globalRand.Int63(), nil
}

func Uint32() (uint32, error) {
	if err != nil {
		return globalRand.Uint32(), err
	}
	return globalRand.Uint32(), nil
}

func Uint64() (uint64, error) {
	if err != nil {
		return globalRand.Uint64(), err
	}
	return globalRand.Uint64(), nil
}

func Int31() (int32, error) {
	if err != nil {
		return globalRand.Int31(), err
	}
	return globalRand.Int31(), nil
}

func Int() (int, error) {
	if err != nil {
		return globalRand.Int(), err
	}
	return globalRand.Int(), nil
}

func Int63n(n int64) (int64, error) {
	if err != nil {
		return globalRand.Int63n(n), err
	}
	return globalRand.Int63n(n), nil
}

func Int31n(n int32) (int32, error) {
	if err != nil {
		return globalRand.Int31n(n), err
	}
	return globalRand.Int31n(n), nil
}

func Intn(n int) (int, error) {
	if err != nil {
		return globalRand.Intn(n), err
	}
	return globalRand.Intn(n), nil
}

func Float64() (float64, error) {
	if err != nil {
		return globalRand.Float64(), err
	}
	return globalRand.Float64(), nil
}

func Float32() (float32, error) {
	if err != nil {
		return globalRand.Float32(), err
	}
	return globalRand.Float32(), nil
}

func Perm(n int) ([]int, error) {
	if err != nil {
		return globalRand.Perm(n), err
	}
	return globalRand.Perm(n), nil
}

func Shuffle(n int, swap func(i, j int)) error {
	if err != nil {
		return err
	}
	globalRand.Shuffle(n, swap)
	return nil
}

func Read(p []byte) (n int, err error) {
	return globalRand.Read(p)
}

func NormFloat64() (float64, error) {
	if err != nil {
		return globalRand.NormFloat64(), err
	}
	return globalRand.NormFloat64(), nil
}

func ExpFloat64() (float64, error) {
	if err != nil {
		return globalRand.ExpFloat64(), err
	}
	return globalRand.ExpFloat64(), nil
}
