package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radiuos float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Circle) Area() float64 {
	return math.Pi * r.Radiuos * r.Radiuos
}
