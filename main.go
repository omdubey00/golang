package main

import (
	"bufio"
	"fmt"
	_ "io"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	fmt.Println("here is f ---- ", f)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

	hello := "game"
}
