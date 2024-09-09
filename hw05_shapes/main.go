package main

import (
	"errors"
	"fmt"

	"github.com/homework-basic/hw05_shapes/figure"
)

func main() {
	shapes := []any{
		figure.NewRectangle(10, 5),
		figure.NewCircle(5),
		figure.NewTriangle(8, 6),
		5,
	}

	for _, shape := range shapes {
		area, err := CalculateArea(shape)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s Площадь: %f\n", shape, area)
	}
}

// CalculateArea ожидает на входе интерфейс Shape и вычисляет площадь
// фигуры, реализующей данный интерфейс.
func CalculateArea(a any) (float32, error) {
	shape, ok := a.(figure.Shape)
	if !ok {
		return 0, errors.New("ошибка: переданный объект не является фигурой")
	}
	return shape.Area(), nil
}
