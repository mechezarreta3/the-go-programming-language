package main

import (
	"fmt"

	"github.com/mechezarreta3/the-go-programming-language/chapter2/examples/popcount/popcount"
)

func main() {
	fmt.Println(popcount.PopCount(15))
	fmt.Println(popcount.PopCountLoop(15))
	fmt.Println(popcount.PopCountShift64(15))
	fmt.Println(popcount.PopCount(1))
	fmt.Println(popcount.PopCountLoop(1))
	fmt.Println(popcount.PopCountShift64(1))
	fmt.Println(popcount.PopCount(7))
	fmt.Println(popcount.PopCountLoop(7))
	fmt.Println(popcount.PopCountShift64(7))
}
