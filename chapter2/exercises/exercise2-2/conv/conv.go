package conv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Celsius temperature to Fahrenheit.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToM converts Feet to Meters.
func FToM(f Feet) Meters { return Meters(f * .3048) }

// MToF converts Meters to Feet.
func MToF(m Meters) Feet { return Feet(m * 3.28084) }

// PToK converts Pounds to Kilograms.
func PToK(p Pounds) Kilograms { return Kilograms(p * 0.45359237) }

// KToP converts Kilograms to Pounds.
func KToP(k Kilograms) Pounds { return Pounds(k * 2.20462) }
