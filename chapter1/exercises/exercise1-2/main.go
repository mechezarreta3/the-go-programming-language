package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Printf("idx: %d, arg: %s\n", idx, arg)
	}
}
