package fib_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	fib "github.com/yyotti/github-actions-test"
)

// TODO CIをコケさせるためにわざとTODOコメントを入れてみる

func test(fn fib.Func, t *testing.T) {
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

func bench(fn fib.Func, args []uint, b *testing.B) {
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
	test(fib.Recursive, t)
}

func BenchmarkRecursive(b *testing.B) {
	bench(fib.Recursive, []uint{10, 20, 30, 40}, b)
}

func TestLoop(t *testing.T) {
	test(fib.Loop, t)
}

func BenchmarkLoop(b *testing.B) {
	bench(fib.Loop, []uint{10, 20, 30, 40, 50, 60, fib.MaxN}, b)
}

func TestGeneralTerm(t *testing.T) {
	test(fib.GeneralTerm, t)
}

func BenchmarkGeneralTerm(b *testing.B) {
	bench(fib.GeneralTerm, []uint{10, 20, 30, 40, 50, 60, fib.MaxN - 21}, b)
}

func TestMapMemoRecursive(t *testing.T) {
	test(fib.MapMemoRecursive, t)
}

func BenchmarkMapMemoRecursive(b *testing.B) {
	bench(fib.MapMemoRecursive, []uint{10, 20, 30, 40, 50, 60, fib.MaxN}, b)
}
