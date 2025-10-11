package structs

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 { // by just watching the syntax of this type of func you can understand it is a method.
	return r.Width * r.Height

	// also here r is a reciever which means whenever this func is called on type Rectangel we will receive all the
	// information related to that type in that referece variable which is by convention the first letter of the type.
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Circle struct {
	Radius float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func Area(rect Rectangle) float64 {
	return (rect.Width * rect.Height)
}

// Here is my understanding that type interface defines an interface on any type and any type that satisfies it
// will pass. Like in above since we were repeating ourselves we turned to build an helper func that takes the
// shape as an arg and checks if Area() can be called on it because Shape is an interface and if we provide a type
// that satisfies it , it will call the Area() func on it . Now satisfaction refers to having the property or the
// method interface is telling and interface resolution is implicit  in go.

// So Interfaces are very broad as compared to types , types are limited to what they have defined while interface
// is valid on any given type if it satisfies it.
