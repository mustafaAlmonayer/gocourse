package intermediate

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perim() float64
}

type Rect struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rect) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius
}

func (r Rect) Perim() float64 {
	return 2 * (r.Height + r.Width)
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) diameter() float64 {
	return 2 * c.Radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func interfaces() {

	r := Rect{Width: 3, Height: 4}
	c := Circle{Radius: 5}

	measure(r)
	measure(c)

}
