// Mandelbrot emits a PNG image of the Mandelbrot fractal with quincunx algorithm super-sampling.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := float64(0); py < height; py++ {
		for px := float64(0); px < width; px++ {
			c := mean([]color.Color{
				mandelbrot(complexPoint(px-0.5, py-0.5)),
				mandelbrot(complexPoint(px-0.5, py+0.5)),
				mandelbrot(complexPoint(px, py)),
				mandelbrot(complexPoint(px+0.5, py-0.5)),
				mandelbrot(complexPoint(px+0.5, py+0.5)),
			})
			img.Set(int(px), int(py), c)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func complexPoint(px, py float64) complex128 {
	y := py/height*(ymax-ymin) + ymin
	x := px/width*(xmax-xmin) + xmin
	return complex(x, y)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			blue := uint8(real(v)*128) + 127
			red := uint8(imag(v)*128) + 127
			return color.YCbCr{255 - contrast*n, blue, red}
		}
	}
	return color.Black
}

func mean(cs []color.Color) color.Color {
	var rs, gs, bs, as uint32
	for _, c := range cs {
		r, g, b, a := c.RGBA()
		rs += r
		gs += g
		bs += b
		as += a
	}

	l := uint32(len(cs))
	rs /= l
	gs /= l
	bs /= l
	as /= l
	return color.RGBA{
		uint8(rs / 0x101),
		uint8(gs / 0x101),
		uint8(bs / 0x101),
		uint8(as / 0x101),
	}
}
