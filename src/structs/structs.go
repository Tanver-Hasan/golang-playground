package structs

import "math"

// Shape interface
type Shape interface{
	Area() float64
}
// Rectangle type
type Rectangle struct {
	Width float64
	Height float64
}
// Circle type
type Circle struct {
	Radius float64
}

// Perimeter - Calculating permeter
func Perimeter(rectangle Rectangle) float64 {
    return 2 * (rectangle.Width + rectangle.Height)
}
// Area - Calculating rectangle area
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Area - calculating circle area
func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

