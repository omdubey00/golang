package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(os.Stdout, "theriddler")

	got := buffer.String()
	want := "theriddler"

	if got != want {
		t.Errorf("got : %q , want : %q ", got, want)
	}
}
