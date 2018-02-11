// make a lissajous GIF
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
	"os"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	if len(os.Args[1:]) > 0 && os.Args[1] == "web" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			cycleStr := r.FormValue("cycles")
			cycles, err := strconv.Atoi(cycleStr)
			if err != nil {
				log.Printf("bad cycles param: %s, %v", cycleStr, err)
				return
			}
			lissajous(w, cycles)
		})
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, 5)
}

func lissajous(out io.Writer, cycles int) {
	const (
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
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
		phase += 0.1
	}
	gif.EncodeAll(out, &anim)
}
