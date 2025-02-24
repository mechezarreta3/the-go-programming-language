package popcount

import "testing"

// pc[i] is the population of count i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Popcount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

func PopCountShift64(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x >>= 1
	}
	return count
}

func PopCountClearNonZero(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func BenchmarkExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(18446744073709551615)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(18446744073709551615)
	}
}

func BenchmarkShift64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift64(18446744073709551615)
	}
}

func BenchmarkClearNonZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearNonZero(18446744073709551615)
	}
}
