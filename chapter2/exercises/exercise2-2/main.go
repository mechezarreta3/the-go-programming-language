package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mechezarreta3/the-go-programming-language/chapter2/exercises/exercise2-2/conv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			value, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid float64 value %v: %v\n", value, err)
				continue
			}
			printConversion(value)
		}

	} else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter float values (press ctrl+d to end): ")
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Fields(line)
			for _, field := range fields {
				value, err := strconv.ParseFloat(field, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Invalid float64 value '%s': %v\n", field, err)
					continue
				}
				printConversion(value)
			}
		}
	}
}

func printConversion(value float64) {
	fmt.Printf("%s = %s,\t%s = %s,\n%s = %s,\t%s = %s,\n%s = %s,\t%s = %s\n\n",
		conv.Fahrenheit(value), conv.FToC(conv.Fahrenheit(value)),
		conv.Celsius(value), conv.CToF(conv.Celsius(value)),
		conv.Feet(value), conv.FToM(conv.Feet(value)),
		conv.Meters(value), conv.MToF(conv.Meters(value)),
		conv.Pounds(value), conv.PToK(conv.Pounds(value)),
		conv.Kilograms(value), conv.KToP(conv.Kilograms(value)))
}
