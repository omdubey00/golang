package structs

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("Perimeter of a rectangle of given widht and height ", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got : %.2f , want : %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct { // we have made some changes into the struct to make it more descriptive regarding the tests.
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, hasArea: 100.0}, // this is called a table driven test in golang
		{name: "Circle", shape: Circle{10.0}, hasArea: 314.1592653589793}, // here we have created a table of tests and iterated through each of
		{name: "Triangle", shape: Triangle{10.0, 10.0}, hasArea: 50.0},    // them and tested each of them , pov they are of the same type .
		// this is making our test more descriptive so that we can understand each field by just seeing them.
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf(" for %#v got %f , want %f ", test.shape, got, test.hasArea)
			}
		})
	} // refactored the test to give each test a testcase using t.Run and was able to make the error message more descriptive.
}
