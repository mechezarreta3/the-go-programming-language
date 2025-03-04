package main

import (
	"fmt"
)

type ByteSize float64

const (
	KB float64 = 1000
	MB float64 = 1000 * KB
	GB float64 = 1000 * MB
	TB float64 = 1000 * GB
	PB float64 = 1000 * TB
	EB float64 = 1000 * PB
	ZB float64 = 1000 * EB
	YB float64 = 1000 * ZB
)

func main() {
	fmt.Printf("KB: %g\nMB: %g\nGB: %g\nTB: %g\nPB: %g\nEB: %g\nZB: %g\nYB: %g\n",
		KB, MB, GB, TB, PB, EB, ZB, YB)
}
