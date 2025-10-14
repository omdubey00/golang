package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.00, 10.00}
	got := Perimeter(rectangle)
	want := 40.00

	if got != want {
		t.Errorf("got : %.2f , want : %.2f ", got, want) // f is the format string for floating point numbers
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got : %g , want : %g ", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		checkArea(t, rectangle, 100.00)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

// so now I have written the bare minimum code to pass the test now refactoring will be done.
// But we have created a general way of testing as there was a clear pattern of taking shapes and finding and comparing their area.
