// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type params struct {
	cycles  int     // number of complete x oscillator revolutions
	size    int     // image canvas covers [-size..+size]
	nframes int     // number of animation frames
	delay   int     // delay between frames in 10ms units
	res     float64 // angular resolution
}

func parseParams(v url.Values) params {
	p := params{
		cycles:  5,
		size:    100,
		nframes: 64,
		delay:   8,
		res:     0.001,
	}
	if c, err := strconv.Atoi(v.Get("cycles")); err == nil {
		p.cycles = c
	}
	if s, err := strconv.Atoi(v.Get("size")); err == nil {
		p.size = s
	}
	if n, err := strconv.Atoi(v.Get("nframes")); err == nil {
		p.nframes = n
	}
	if d, err := strconv.Atoi(v.Get("delay")); err == nil {
		p.delay = d
	}
	if r, err := strconv.ParseFloat(v.Get("res"), 64); err == nil {
		p.res = r
	}
	return p
}

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			lissajous(w, parseParams(r.Form))
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, parseParams(nil))
}

func lissajous(out io.Writer, p params) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: p.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < p.nframes; i++ {
		rect := image.Rect(0, 0, 2*p.size+1, 2*p.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.cycles)*2*math.Pi; t += p.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(p.size+int(x*float64(p.size)+0.5), p.size+int(y*float64(p.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
