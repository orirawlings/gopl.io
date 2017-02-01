package weightconv

import "fmt"

type Pound float64
type Gram float64

func (p Pound) String() string { return fmt.Sprintf("%g lb", p) }
func (g Gram) String() string  { return fmt.Sprintf("%g g", g) }

// GToLb converts Gram weight to Pound.
func GToLb(g Gram) Pound { return Pound(g * 0.0022046) }

// LbToG converts Pound weight to Gram.
func LbToG(p Pound) Gram { return Gram(p / 0.0022046) }
