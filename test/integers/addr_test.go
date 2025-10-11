package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("got : %q , want : %q", got, want)
	}
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}
