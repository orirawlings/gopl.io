// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var (
	sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
	top          = color{0, 0, 0xff}                // color for top z values
	bottom       = color{0xff, 0, 0}                // color for bottom z values
)

type color struct {
	r, g, b uint8
}

func (c color) String() string {
	return fmt.Sprintf("#%02x%02x%02x", c.r, c.g, c.b)
}

type value struct {
	x, y, z float64
	ok      bool
}

func main() {
	var values [cells + 1][cells + 1]value
	zmin, zmax := math.Inf(1), math.Inf(-1)
	for i := 0; i <= cells; i++ {
		for j := 0; j <= cells; j++ {
			x, y, z, ok := corner(i, j)
			if ok {
				zmin, zmax = math.Min(z, zmin), math.Max(z, zmax)
			}
			values[i][j] = value{x, y, z, ok}
		}
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			a := values[i+1][j]
			b := values[i][j]
			c := values[i][j+1]
			d := values[i+1][j+1]
			if a.ok && b.ok && c.ok && d.ok {
				blend := (b.z - zmin) / (zmax - zmin)
				fmt.Printf("<polygon style='fill: %v' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					clerp(bottom, top, blend), a.x, a.y, b.x, b.y, c.x, c.y, d.x, d.y)
			}
		}
	}
	fmt.Println("</svg>")
}

// Linear Interpolation between v0 and v1
func lerp(v0, v1, t float64) float64 {
	return (1-t)*v0 + t*v1
}

// Linear Interpolation between colors v0 and v1
func clerp(v0, v1 color, t float64) color {
	return color{
		uint8(lerp(float64(v0.r), float64(v1.r), t)),
		uint8(lerp(float64(v0.g), float64(v1.g), t)),
		uint8(lerp(float64(v0.b), float64(v1.b), t)),
	}
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return 0, 0, z, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

// Returns value and true if value is finite, 0 and false otherwise
func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	if r == 0 {
		return 0, false
	}
	return math.Sin(r) / r, true
}
