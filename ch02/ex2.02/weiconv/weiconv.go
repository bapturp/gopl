package weiconv

import "fmt"

type Kilogram float64
type Pound float64

// String representation of Kilogram weight
func (k Kilogram) String() string { return fmt.Sprintf("%g Kilogram", k) }

// String representation of Pound weight
func (p Pound) String() string { return fmt.Sprintf("%g Pound", p) }

// KToP converts a Kilogram weight in Pound
func KToP(k Kilogram) Pound { return Pound(k / 0.45359237) }

// PToK converts a Pound weight in Kilogram
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }
