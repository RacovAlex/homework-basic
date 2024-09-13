package service

import (
	"fmt"
	"testing"

	"github.com/homework-basic/hw06_testing/figure"
	"github.com/stretchr/testify/assert"
)

type MockShape struct {
	area float32
}

func (m *MockShape) Area() (float32, error) {
	return m.area, nil
}

func (m *MockShape) String() string {
	return fmt.Sprintf("area:%f", m.area)
}

func TestCalc_Area(t *testing.T) {
	const expected float32 = 10
	calc := Calc{}
	shapes := []figure.Shape{
		&MockShape{expected},
	}
	got := calc.Area(shapes...)
	assert.Equal(t, expected, got)
}
