package figure

import "fmt"

// Triangle - фигура треугольника с основанием и высотой.
type Triangle struct {
	base, height float32
}

func NewTriangle(base, height float32) Shape {
	return &Triangle{
		base:   base,
		height: height,
	}
}

// Area - метод вычисляющий площадь треугольника.
func (t *Triangle) Area() (float32, error) {
	return 0.5 * t.base * t.height, nil
}

// String - выводит описание треугольника.
func (t *Triangle) String() string {
	return fmt.Sprintf("Треугольник: основание %f, высота %f", t.base, t.height)
}
