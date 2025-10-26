package main

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, text string) {
	fmt.Fprintf(w, "%s", text)
}
