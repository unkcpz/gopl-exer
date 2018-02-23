// popcount test
package popcount_test

import (
	"testing"

	"github.com/unkcpz/gopl-exer/ch2/popcount/popcount"
)

func TestPopCount(t *testing.T) {
	if !(popcount.PopCount(0) == 0) {
		t.Errorf("popcount.PopCount(0) = %d\n", popcount.PopCount(0))
	}
}
