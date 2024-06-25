package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "Hello, World!\n"
	output := captureOutput(main)
	if output != expected {
		t.Errorf("Expected %q, but got %q", expected, output)
	}
}

// captureOutput captures output of a function
func captureOutput(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old // restoring the real stdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
