// Package tempconv performs Celsius and Fahrenheit conversions.
package conv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC = -273.15
	FreezingC     = 0
	BoilingC      = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%.1f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.1f°F", f) }
