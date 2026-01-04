package intermediate

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type MyInt int

func (m MyInt) IsPositive() bool {
	return m > 0
}

func methods() {
	rect := Rectangle{
		length: 10,
		width:  9,
	}
	area := rect.Area()
	fmt.Println(area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println(area)

	num := MyInt(5)
	fmt.Println(num.IsPositive())

	s := Shape{
		Rectangle: Rectangle{
			length: 10,
			width:  6,
		},
	}

	s.Scale(2)

}
