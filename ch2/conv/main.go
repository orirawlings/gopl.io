// conv converts its numeric argument to values in various units
package main

import (
	"fmt"
	"gopl.io/ch2/lengthconv"
	"gopl.io/ch2/tempconv1"
	"gopl.io/ch2/weightconv"
	"os"
	"strconv"
)

func temp(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	k := tempconv.Kelvin(t)
	fmt.Printf("%s = %s = %s\n%s = %s = %s\n%s = %s = %s\n",
		f, tempconv.FToC(f), tempconv.FToK(f),
		c, tempconv.CToF(c), tempconv.CToK(c),
		k, tempconv.KToC(k), tempconv.KToF(k))
}

func length(l float64) {
	f := lengthconv.Feet(l)
	m := lengthconv.Meter(l)
	fmt.Printf("%s = %s\n%s = %s\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
}

func weight(w float64) {
	p := weightconv.Pound(w)
	g := weightconv.Gram(w)
	fmt.Printf("%s = %s\n%s = %s\n", p, weightconv.LbToG(p), g, weightconv.GToLb(g))
}

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		temp(v)
		length(v)
		weight(v)
	}
}
