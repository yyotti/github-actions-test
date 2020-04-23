package fib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckN(t *testing.T) {
	a := assert.New(t)

	for n := uint(0); n <= MaxN; n++ {
		n := n
		a.NotPanics(func() { checkN(n) }, "n={}", n)
	}

	a.Panics(func() { checkN(MaxN + 1) }, "n={}", MaxN+1)
}

func test(fn Func, t *testing.T) {
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

func TestMapMemoRecursive_inner(t *testing.T) {
	test(func(n uint) uint64 {
		memo := map[uint]uint64{}
		return mapMemoRecursive(n, memo)
	}, t)
}

func TestArrayMemoRecursive_inner(t *testing.T) {
	test(func(n uint) uint64 {
		memo := make([]*uint64, n+1)
		return arrayMemoRecursive(n, memo)
	}, t)
}
