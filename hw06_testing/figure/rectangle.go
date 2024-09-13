package figure

import (
	"errors"
	"fmt"
)

var ErrNegativeHeightWidth = errors.New("высота или ширина не может быть отрицательной")

// Rectangle - фигура прямоугольника с шириной и высотой.
type Rectangle struct {
	width, height float32
}

func NewRectangle(width, height float32) Shape {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

// Area - метод вычисляющий площадь прямоугольника.
func (r *Rectangle) Area() (float32, error) {
	if r.width < 0 || r.height < 0 {
		return 0, ErrNegativeHeightWidth
	}
	return r.width * r.height, nil
}

// String - выводит описание прямоугольника.
func (r *Rectangle) String() string {
	return fmt.Sprintf("Прямоугольник: ширина %f, высота %f", r.width, r.height)
}
