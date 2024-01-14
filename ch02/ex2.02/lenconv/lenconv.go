package lenconv

import "fmt"

type Meter float64
type Foot float64

// String representation of a Meter length
func (m Meter) String() string { return fmt.Sprintf("%g Meter", m) }

// String representation of Foot length
func (f Foot) String() string { return fmt.Sprintf("%g Foot", f) }

// MToF converts a Meter length in Foot
func MToF(m Meter) Foot { return Foot(m * 3.28084) }

// FToM converts a Foot length in Meter
func FToM(f Foot) Meter { return Meter(f * 0.3048) }
