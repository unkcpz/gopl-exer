package geometry

import "testing"

var p, q Point = Point{1, 3}, Point{4, 7}
var wanted float64 = 5

func TestDistance(t *testing.T) {
	if got := Distance(p, q); got != wanted {
		t.Errorf("Distance(%#v, %#v) = %f, want %f", p, q, got, wanted)
	}
}

func TestPointDistance(t *testing.T) {
	if got := p.Distance(q); got != wanted {
		t.Errorf("%#v.Distance(%#v) = %f, want %f", p, q, got, wanted)
	}
}

func TestPathDistance(t *testing.T) {
	var path Path = Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	var wanted float64 = 12.0
	if got := path.Distance(); got != wanted {
		t.Errorf("%#v.Distance() = %f, want %f", path, got, wanted)
	}
}
