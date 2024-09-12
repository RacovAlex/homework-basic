package figure

import "fmt"

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
	return r.width * r.height, nil
}

// String - выводит описание прямоугольника.
func (r *Rectangle) String() string {
	return fmt.Sprintf("Прямоугольник: ширина %f, высота %f", r.width, r.height)
}
