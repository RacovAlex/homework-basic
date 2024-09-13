package service

import "github.com/homework-basic/hw06_testing/figure"

type Calc struct{}

func (c *Calc) Area(shapes ...figure.Shape) float32 {
	var results float32
	for _, shape := range shapes {
		result, _ := shape.Area()
		results += result
	}
	return results
}
