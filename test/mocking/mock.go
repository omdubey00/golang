package main

import (
	"fmt"
	"io"
	"os"
)

func CountDown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func main() {
	CountDown(os.Stdout)
}

func Gamer() {
	const hello = "gamer"
}

// hopping between multiple files is so easy with harpoon and it's the single biggest reason I love
// nvim is navigation is so fast
// I can just  go to the my primary file with one keypress and using harpoon I just showed you can hop
// between files. Let me show you how I move from one to 5 files faxst  ....
