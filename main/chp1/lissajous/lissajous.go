package main

//Lissajous generate GIF animations of random Lissajous figures.

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 //first color in palette
	blackIndex = 1 //nextColor in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //Note: ignoring encoding errors
}

/*
Variable color/white belongs to the image/color package and gif.GIF belongs to image/gif

We refer a package whose path has const declaration that gives names to constants that are values that are fiexed
at compile time here.

color.Colot{...} and gif.GIF{...} are composite literals
a compact notation for instantiating any of Go's composite types from a sequence of element values.

First one is a slice and a second one is a struct.

gif.GIF is a struct type. A struct is a group of values called
fields, often of different tyeps that are collected together in a single object.

The struct literatl creates a struct value whose Loop-Count field is set to nframes;
All other fields have zero value for their type





*/
