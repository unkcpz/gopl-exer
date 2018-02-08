// make a lissajous GIF
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var colorGreen = color.RGBA{0x33, 0xcc, 0x33, 0xff}
var colorYello = color.RGBA{0xff, 0xff, 0x00, 0xff}
var colorRed = color.RGBA{0xff, 0x00, 0x00, 0xff}
var palette = []color.Color{color.White, colorGreen,
	colorYello, colorRed}

// const (
// 	whiteIndex = 0
// 	blackIndex = 1
// )

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		nframes = 64
		delay   = 8   //delay between frames in 10ms units
		size    = 100 // canvas size
	)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	freq := rand.Float64()
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			idx := uint(x*size) % 3
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(idx))
		}
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
		phase += 0.1
	}
	gif.EncodeAll(out, &anim)
}
