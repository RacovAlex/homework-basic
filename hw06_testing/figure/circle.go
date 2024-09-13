package figure

import (
	"errors"
	"fmt"
	"math"
)

var ErrNegativeRadius = errors.New("радиус не может быть отрицательным")

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
func (c *Circle) Area() (float32, error) {
	if c.radius < 0 {
		return 0.0, ErrNegativeRadius
	}
	return math.Pi * c.radius * c.radius, nil
}

// String - выводит описание круга.
func (c *Circle) String() string {
	return fmt.Sprintf("Круг: радиус %f", c.radius)
}
