// Testing and benchmarking three echo function
package main

import (
	"testing"
)

var testArgs = []string{"cmdn", "arg1", "arg2", "arg3", "arg4", "arg5"}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(testArgs)
	}
}

func BenchmarkRangeConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(testArgs)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(testArgs)
	}
}
