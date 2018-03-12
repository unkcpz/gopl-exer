// geometry used for geometry calculation
package geometry

import "math"

type Point struct{ X, Y float64 }

type Path []Point

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Path) Distance() float64 {
	sum := 0.0
	for i, _ := range p[:len(p)-1] {
		sum += p[i].Distance(p[i+1])
	}
	return sum
}
