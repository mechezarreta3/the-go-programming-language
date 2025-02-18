package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string

	// Method 1: Traditional for loop
	start_time := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("Total Runtime | Method 1: ", time.Since(start_time).Nanoseconds(), "nanoseconds")

	// Method 2: Range for loop
	start_time = time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("Total Runtime | Method 2: ", time.Since(start_time).Nanoseconds(), "nanoseconds")

	// Method 3: strings Join
	start_time = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println("Total Runtime | Method 3: ", time.Since(start_time).Nanoseconds(), "nanoseconds")
}
