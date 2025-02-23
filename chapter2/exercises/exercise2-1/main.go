package main

import "fmt"
import "github.com/mechezarreta3/the-go-programming-language/chapter2/exercises/exercise2-1/tempconv"

func main() {
	fmt.Printf("273.15°K = %v\n", tempconv.KToC(273.15))
	fmt.Printf("-273.15°C = %v\n", tempconv.CToK(-273.15))
}
