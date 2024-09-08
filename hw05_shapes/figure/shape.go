package figure

// Shape интерфейс описывает фигуры, у которых можно вычесть их площадь.
type Shape interface {
	Area() float32
	String() string
}
