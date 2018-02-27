// Surface computes an SVG redering of a 3-D surface function
package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	usage := "usage: ex3_2 saddle | drip | eggbox"
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	var f zFunc
	switch os.Args[1] {
	case "saddle":
		f = saddle
	case "drip":
		f = drip
	case "eggbox":
		f = eggbox
	default:
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	svgPlotter(os.Stdout, f)
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func drip(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

type zFunc func(x, y float64) float64

func svgPlotter(w io.Writer, f zFunc) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i int, j int, f zFunc) (sx float64, sy float64) {
	// Find point (x,y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)

	// Project (x,y,z) ontu 2-D
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
