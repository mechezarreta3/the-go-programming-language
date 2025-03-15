package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice X Len: %d, Slice X Cap: %d\nSlice X: %v\n", len(x), cap(x), x)
	x = appendInt(x, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("Slice X Len: %d, Slice X Cap: %d\nSlice X: %v\n", len(x), cap(x), x)

	var a, b []int
	for i := 0; i < 10; i++ {
		b = appendInt(a, i, i)
		fmt.Printf("%-2d Cap=%-4d%v\n", i, cap(b), b)
		a = b
	}
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text)
	}

	copy(z[len(x):], y)
	return z
}
