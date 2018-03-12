// variadic functions max and min
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(max(3.0, -9, 9))  // 9
	fmt.Println(max0(3.0, -9, 9)) // 9
	fmt.Println(min(3, -9.1, 9))  // -9.1
	fmt.Println(max())            // NaN
}

func max(vals ...float64) float64 {
	if len(vals) == 0 {
		return math.NaN()
	}
	m := vals[0]
	for _, val := range vals[1:] {
		if val > m {
			m = val
		}
	}
	return m
}

func min(vals ...float64) float64 {
	if len(vals) == 0 {
		return math.NaN()
	}
	m := vals[0]
	for _, val := range vals[1:] {
		if val < m {
			m = val
		}
	}
	return m
}

// at least one variable
func max0(v float64, vals ...float64) float64 {
	m := v
	for _, val := range vals {
		if val > m {
			m = val
		}
	}
	return m
}
