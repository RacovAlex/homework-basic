package figures

import "math"

// Rectangle - фигура прямоугольника с шириной и высотой.
type Rectangle struct {
	Width, Height float32
}

func NewRectangle(width, height float32) *Rectangle {
	return &Rectangle{width, height}
}

// Area - метод вычисляющий площадь прямоугольника.
func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

// Circle - фигура круга с радиусом.
type Circle struct {
	Radius float32
}

func NewCircle(radius float32) *Circle {
	return &Circle{radius}
}

// Area - метод вычисляющий площадь круга.
func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle - фигура треугольника с основанием и высотой
type Triangle struct {
	Base, Height float32
}

func NewTriangle(base, height float32) *Triangle {
	return &Triangle{base, height}
}

// Area - метод вычисляющий площадь треугольника.
func (t Triangle) Area() float32 {
	return 0.5 * t.Base * t.Height
}
