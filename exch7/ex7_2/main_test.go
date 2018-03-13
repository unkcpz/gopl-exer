package main

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	b := &bytes.Buffer{}
	_, n := CountingWriter(b)
	data := []byte("hi there")
	if *n != int64(len(data)) {
		t.Logf("%d != %d", n, len(data))
		t.Fail()
	}
}
