// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// String prints its value.
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// String prints its value.
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// String prints its value.
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius(f-32) * 5 / 9 }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k) - AbsoluteZeroC }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + +AbsoluteZeroC) }
