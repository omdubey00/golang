package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello , %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "theriddler")
}

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(MyGreetHandler))
}
