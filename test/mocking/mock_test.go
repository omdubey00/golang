package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	CountDown(buffer)

	got := buffer.String()
	want := "3"
	if got != want {
		t.Errorf("got : %q , want : %q ", got, want)
	}
}

// Just started using obs will be sharing a lot on this account
// until then here is a little peak of my setup or workflow in nvim.
