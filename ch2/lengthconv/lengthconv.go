package lengthconv

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }

// MToF converts Meter length to Feet.
func MToF(m Meter) Feet { return Feet(m * 3.2808) }

// FToM converts Feet length to Meter.
func FToM(f Feet) Meter { return Meter(f / 3.2808) }
