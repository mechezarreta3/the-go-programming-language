// Package tempconv performs Celsius and Fahrenheit conversions.
package conv

import "fmt"

type Feet float64
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%.1fft", f) }
func (m Meters) String() string { return fmt.Sprintf("%.1fm", m) }
