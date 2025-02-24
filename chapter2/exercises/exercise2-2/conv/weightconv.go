// Package tempconv performs Celsius and Fahrenheit conversions.
package conv

import "fmt"

type Pounds float64
type Kilograms float64

func (p Pounds) String() string    { return fmt.Sprintf("%.1flbs", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%.1fkg", k) }
