// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	cx, cy, zoom           = 0.275 + 1.0/(1<<7) + 1.0/(1<<8) + 1.0/(1<<9), 1.0/(1<<6) - 1.0/(1<<13) - 1.0/(1<<14) + 1.0/(1<<16), 1.0 / (1 << 15)
	xmin, ymin, xmax, ymax = cx - zoom, cy - zoom, cx + zoom, cy + zoom
	width, height          = 1024, 1024
	iterations             = 200
	contrast               = 15
)

func main() {
	drawMandelbrotFloat64(os.Stdout)
	// drawMandelbrotComplex64(os.Stdout)
	// drawMandelbrotComplex128(os.Stdout)
	// drawMandelbrotBigFloat(os.Stdout)
	// drawMandelbrotBigRat(os.Stdout)
}

func drawMandelbrotFloat64(w io.Writer) {
	escapes := func(x, y float64) bool {
		h2 := x*x + y*y
		return h2 > 4 || h2 < -4
	}

	mandelbrot := func(r, i float64) color.Color {
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

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func drawMandelbrotComplex64(w io.Writer) {
	mandelbrot := func(z complex64) color.Color {
		var v complex64
		for n := uint8(0); n < iterations; n++ {
			v = v*v + z
			if cmplx.Abs(complex128(v)) > 2 {
				blue := uint8(255 * (real(z) - xmin) / (xmax - xmin))
				red := uint8(255 * (imag(z) - ymin) / (ymax - ymin))
				return color.YCbCr{255 - contrast*n, blue, red}
			}
		}
		return color.Black
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func drawMandelbrotComplex128(w io.Writer) {
	mandelbrot := func(z complex128) color.Color {
		var v complex128
		for n := uint8(0); n < iterations; n++ {
			v = v*v + z
			if cmplx.Abs(v) > 2 {
				blue := uint8(255 * (real(z) - xmin) / (xmax - xmin))
				red := uint8(255 * (imag(z) - ymin) / (ymax - ymin))
				return color.YCbCr{255 - contrast*n, blue, red}
			}
		}
		return color.Black
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func drawMandelbrotBigFloat(w io.Writer) {
	x, y, x2, y2, t := new(big.Float), new(big.Float), new(big.Float), new(big.Float), new(big.Float)

	two := new(big.Float).SetInt64(2)
	four := new(big.Float).SetInt64(4)
	negFour := new(big.Float).SetInt64(-4)

	escapes := func(x, y *big.Float) bool {
		x2.Mul(x, x)
		y2.Mul(y, y)
		t.Add(x2, y2)
		return t.Cmp(four) > 0 || t.Cmp(negFour) < 0
	}

	mandelbrot := func(r, i *big.Float) color.Color {
		x.Set(r)
		y.Set(i)
		for n := uint8(0); n < iterations; n++ {
			t.Set(x)

			// x_n = x_{n-1}^2 - y_{n-1}^2 + x_0
			x.Sub(x2.Mul(x, x), y2.Mul(y, y))
			x.Add(x, r)

			// y_n = 2*x_{n-1}*y_{n-1} + y_0
			y.Mul(y, two)
			y.Mul(y, t)
			y.Add(y, i)

			if escapes(x, y) {
				rf, _ := r.Float64()
				imagf, _ := i.Float64()
				blue := uint8(255 * (rf - xmin) / (xmax - xmin))
				red := uint8(255 * (imagf - ymin) / (ymax - ymin))
				return color.YCbCr{255 - contrast*n, blue, red}
			}
		}
		return color.Black
	}

	xscale := new(big.Float).Quo(big.NewFloat(xmax-xmin), new(big.Float).SetInt64(width))
	yscale := new(big.Float).Quo(big.NewFloat(ymax-ymin), new(big.Float).SetInt64(height))
	xminF, yminF := big.NewFloat(xmin), big.NewFloat(ymin)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	r, i := new(big.Float), new(big.Float)
	for py := 0; py < height; py++ {
		i.SetInt64(int64(py))
		i.Mul(i, yscale)
		i.Add(i, yminF)
		for px := 0; px < width; px++ {
			r.SetInt64(int64(px))
			r.Mul(r, xscale)
			r.Add(r, xminF)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(r, i))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func drawMandelbrotBigRat(w io.Writer) {
	x, y, x2, y2, t := new(big.Rat), new(big.Rat), new(big.Rat), new(big.Rat), new(big.Rat)

	two := new(big.Rat).SetInt64(2)
	four := new(big.Rat).SetInt64(4)
	negFour := new(big.Rat).SetInt64(-4)

	escapes := func(x, y *big.Rat) bool {
		x2.Mul(x, x)
		y2.Mul(y, y)
		t.Add(x2, y2)
		return t.Cmp(four) > 0 || t.Cmp(negFour) < 0
	}

	mandelbrot := func(r, i *big.Rat) color.Color {
		x.Set(r)
		y.Set(i)
		for n := uint8(0); n < iterations; n++ {
			t.Set(x)

			// x_n = x_{n-1}^2 - y_{n-1}^2 + x_0
			x.Sub(x2.Mul(x, x), y2.Mul(y, y))
			x.Add(x, r)

			// y_n = 2*x_{n-1}*y_{n-1} + y_0
			y.Mul(y, two)
			y.Mul(y, t)
			y.Add(y, i)

			if escapes(x, y) {
				rf, _ := r.Float64()
				imagf, _ := i.Float64()
				blue := uint8(255 * (rf - xmin) / (xmax - xmin))
				red := uint8(255 * (imagf - ymin) / (ymax - ymin))
				return color.YCbCr{255 - contrast*n, blue, red}
			}
		}
		return color.Black
	}

	xscale := new(big.Rat).Quo(new(big.Rat).SetFloat64(xmax-xmin), new(big.Rat).SetInt64(width))
	yscale := new(big.Rat).Quo(new(big.Rat).SetFloat64(ymax-ymin), new(big.Rat).SetInt64(height))
	xminF, yminF := new(big.Rat).SetFloat64(xmin), new(big.Rat).SetFloat64(ymin)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	r, i := new(big.Rat), new(big.Rat)
	for py := 0; py < height; py++ {
		i.SetInt64(int64(py))
		i.Mul(i, yscale)
		i.Add(i, yminF)
		for px := 0; px < width; px++ {
			r.SetInt64(int64(px))
			r.Mul(r, xscale)
			r.Add(r, xminF)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(r, i))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}
