package figure

import (
	"fmt"
	"math"
)

// Circle - фигура круга с радиусом.
type Circle struct {
	radius float32
}

func NewCircle(radius float32) Shape {
	return &Circle{
		radius: radius,
	}
}

// Area - метод вычисляющий площадь круга.
func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

// String - выводит описание круга.
func (c *Circle) String() string {
	return fmt.Sprintf("Круг: радиус %f", c.radius)
}
