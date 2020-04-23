// Package fib contains functions for calculating Fibonacci numbers.
// These are all experimental.
package fib

import (
	"fmt"
	"math"
)

type Func func(uint) uint64

// MaxN is limit value of index n.
// fib(94) =  1293530146158671551
// fib(95) = 13493690561280548289
// fib(96) = fib(95) + fib(94) = 14787220707439219840
// fib(97) = fib(96) + fib(95) = 28280911269719768129 > 18446744073709551615
const MaxN uint = 96

func checkN(n uint) {
	if n > MaxN {
		panic(fmt.Sprintf("too large N: %d (> %d)", n, MaxN))
	}
}

func Recursive(n uint) uint64 {
	checkN(n)

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Recursive(n-1) + Recursive(n-2)
	}
}

func Loop(n uint) uint64 {
	checkN(n)

	a, b := uint64(0), uint64(1)
	for i := uint(0); i < n; i++ {
		a, b = b, a+b
	}

	return a
}

func GeneralTerm(n uint) uint64 {
	if n > MaxN-21 {
		panic(fmt.Sprintf("too large N: %d (> %d)", n, MaxN-21))
	}

	// phi = (1 + sqrt(5)) / 2
	// F(n) = floor(phi^n / sqrt(5) + 1/2)
	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2

	return uint64(math.Floor(math.Pow(phi, float64(n))/sqrt5 + 1.0/2))
}

func mapMemoRecursive(n uint, memo map[uint]uint64) uint64 {
	if fib, ok := memo[n]; ok {
		return fib
	}

	switch n {
	case 0:
		memo[0] = 0
		return 0
	case 1:
		memo[1] = 1
		return 1
	default:
		a := mapMemoRecursive(n-1, memo)
		b := mapMemoRecursive(n-2, memo)
		memo[n] = a + b
	}

	return mapMemoRecursive(n, memo)
}

func MapMemoRecursive(n uint) uint64 {
	checkN(n)

	return mapMemoRecursive(n, make(map[uint]uint64, n))
}

func arrayMemoRecursive(n uint, memo []*uint64) uint64 {
	if fib := memo[n]; fib != nil {
		return *fib
	}

	var fib uint64

	switch n {
	case 0:
		fib = 0
	case 1:
		fib = 1
	default:
		a := arrayMemoRecursive(n-1, memo)
		b := arrayMemoRecursive(n-2, memo)
		fib = a + b
	}

	memo[n] = &fib

	return arrayMemoRecursive(n, memo)
}

func ArrayMemoRecursive(n uint) uint64 {
	checkN(n)

	return arrayMemoRecursive(n, make([]*uint64, n+1))
}
