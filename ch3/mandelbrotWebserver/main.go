// Mandelbrot webserver that returns a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 1024, 1024
	iterations    = 200
	contrast      = 15
)

func param(form map[string][]string, p string, d float64) float64 {
	if vs, ok := form[p]; ok {
		for _, v := range vs {
			if r, err := strconv.ParseFloat(v, 64); err == nil {
				return r
			}
		}
	}
	return d
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		r.ParseForm()
		x := param(r.Form, "x", 0)
		y := param(r.Form, "y", 0)
		zoom := param(r.Form, "zoom", -1)
		drawMandelbrotFloat64(w, x, y, 1.0/(math.Pow(2, zoom)))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func escapes(x, y float64) bool {
	h2 := x*x + y*y
	return h2 > 4 || h2 < -4
}

func mandelbrot(r, i, xmin, xmax, ymin, ymax float64) color.Color {
	var x, y float64
	for n := uint8(0); n < iterations; n++ {
		x, y = x*x-y*y+r, 2*x*y+i
		if escapes(x, y) {
			blue := uint8(255 * (r - xmin) / (xmax - xmin))
			red := uint8(255 * (i - ymin) / (ymax - ymin))
			return color.YCbCr{255 - contrast*n, blue, red}
		}
	}
	return color.Black
}

func drawMandelbrotFloat64(w io.Writer, cx, cy, r float64) {
	xmin, ymin, xmax, ymax := cx-r, cy-r, cx+r, cy+r
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(x, y, xmin, xmax, ymin, ymax))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}
