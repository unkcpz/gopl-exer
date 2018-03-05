// embed use struct embed syntax sugar
package main

import "fmt"

type Points struct {
	X, Y int
}

type Circle struct {
	Points
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{
		Circle: Circle{
			Points: Points{5, 8},
			Radius: 10,
		},
		Spokes: 2,
	}

	fmt.Printf("%#v\n", w)

	w.X = 99

	fmt.Printf("%#v\n", w)
}
