package fib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckN(t *testing.T) {
	a := assert.New(t)
	for n := uint(0); n <= MaxN; n++ {
		a.NotPanics(func() { checkN(n) }, "n={}", n)
	}

	a.Panics(func() { checkN(MaxN + 1) }, "n={}", MaxN+1)
}

func test(fn FibFunc, t *testing.T) {
	var tests = []struct {
		name     string
		expected uint64
		given    uint
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"2", 1, 2},
		{"3", 2, 3},
		{"4", 3, 4},
		{"5", 5, 5},
		{"6", 8, 6},
		{"7", 13, 7},
		{"8", 21, 8},
		{"9", 34, 9},
		{"10", 55, 10},
	}

	a := assert.New(t)
	for _, tt := range tests {
		a.Equal(tt.expected, fn(tt.given), "n=%d", tt.given)
	}
}

func bench(fn FibFunc, args []uint, b *testing.B) {
	for _, arg := range args {
		arg := arg
		b.Run(fmt.Sprintf("%d", arg), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				fn(arg)
			}
		})
	}
}

func TestRecursive(t *testing.T) {
	test(Recursive, t)
}

func BenchmarkRecursive(b *testing.B) {
	bench(Recursive, []uint{10, 20, 30, 40}, b)
}

func TestLoop(t *testing.T) {
	test(Loop, t)
}

func BenchmarkLoop(b *testing.B) {
	bench(Loop, []uint{10, 20, 30, 40, 50, 60, MaxN}, b)
}

func TestGeneralTerm(t *testing.T) {
	test(GeneralTerm, t)
}

func BenchmarkGeneralTerm(b *testing.B) {
	bench(GeneralTerm, []uint{10, 20, 30, 40, 50, 60, MaxN - 21}, b)
}

func TestMapMemoRecursive_inner(t *testing.T) {
	test(func(n uint) uint64 {
		memo := map[uint]uint64{}
		return mapMemoRecursive(n, memo)
	}, t)
}

func TestMapMemoRecursive(t *testing.T) {
	test(MapMemoRecursive, t)
}

func BenchmarkMapMemoRecursive(b *testing.B) {
	bench(MapMemoRecursive, []uint{10, 20, 30, 40, 50, 60, MaxN}, b)
}

func TestArrayMemoRecursive_inner(t *testing.T) {
	test(func(n uint) uint64 {
		memo := make([]*uint64, n+1)
		return arrayMemoRecursive(n, memo)
	}, t)
}

func TestArrayMemoRecursive(t *testing.T) {
	test(ArrayMemoRecursive, t)
}

func BenchmarkArrayMemoRecursive(b *testing.B) {
	bench(ArrayMemoRecursive, []uint{10, 20, 30, 40, 50, 60, MaxN}, b)
}
