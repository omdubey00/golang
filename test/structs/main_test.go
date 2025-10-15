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
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 10, height: 10}, want: 100.00},
		{name: "Rectangle", shape: Circle{Radius: 10}, want: 314.1592653589793}, // This is we have created a list of test of same type.
		{name: "Triangle", shape: Triangle{10, 2}, want: 10.00},
	}

	for _, ts := range areaTests {
		t.Run(ts.name, func(t *testing.T) {
			got := ts.shape.Area()

			if got != ts.want {
				t.Errorf("shape : %#v ,got : %g , want : %g ", ts.shape, got, ts.want)
			}
		})
	}
}

// so now I have written the bare minimum code to pass the test now refactoring will be done.
// But we have created a general way of testing as there was a clear pattern of taking shapes and finding and comparing their area.
