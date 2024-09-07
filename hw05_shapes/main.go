package main

import (
	"errors"
	"fmt"

	"github.com/homework-basic/hw05_shapes/figures"
)

// Shape интерфейс описывает фигуры, у которых можно вычесть их площадь.
type Shape interface {
	Area() float32
}

func main() {
	PrintFigureWithArea(5)
}

// CalculateArea ожидает на входе интерфейс Shape и вычисляет площадь
// фигуры, реализующей данный интерфейс.
func CalculateArea(a any) (float32, error) {
	figure, ok := a.(Shape)
	if !ok {
		return 0, errors.New("ошибка: переданный объект не является фигурой")
	}
	return figure.Area(), nil
}

func PrintFigureWithArea(a any) {
	area, err := CalculateArea(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch a := a.(type) {
	case *figures.Circle:
		fmt.Printf("Круг: радиус %f Площадь: %f\n", a.Radius, area)
	case *figures.Rectangle:
		fmt.Printf("Прямоугольник: ширина %f, высота %f Площадь: %f\n", a.Width, a.Height, area)
	case *figures.Triangle:
		fmt.Printf("Треугольник: основание %f, высота %f Площадь: %f\n", a.Base, a.Height, area)
	}
}
