package visitor

import (
	"fmt"
	"math"
)

type Visitor interface { // 访问者抽象.  为每种要访问的结构定义访问方法.
	VisitForSquare(shape *Square)
	VisitForCircle(shape *Circle)
	VisitForRectangle(shape *Rectangle)
}

type AreaCalculator struct { // 面积访问者
	area float64
}

func (a *AreaCalculator) VisitForSquare(s *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	a.area = float64(s.Side * s.Side)
	fmt.Printf("Calculating area for square:%.2f\n", a.area)
}

func (a *AreaCalculator) VisitForCircle(s *Circle) {
	a.area = float64(2) * math.Pi * float64(s.Radius*s.Radius)
	fmt.Printf("Calculating area for circle:%.2f\n", a.area)
}
func (a *AreaCalculator) VisitForRectangle(s *Rectangle) {
	a.area = float64(s.L * s.B)
	fmt.Printf("Calculating area for rectangle:%.0f\n", a.area)
}

type MiddleCoordinates struct {
	x float64
	y float64
}

func (a *MiddleCoordinates) VisitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	a.x = 0.5 * float64(s.Side)
	a.y = 0.5 * float64(s.Side)
	fmt.Printf("Calculating middle point coordinates for square:%+v\n", a)
}

func (a *MiddleCoordinates) VisitForCircle(c *Circle) {
	a.x = float64(c.Radius)
	a.y = float64(c.Radius)
	fmt.Printf("Calculating middle point coordinates for circle:%+v\n", a)
}
func (a *MiddleCoordinates) VisitForRectangle(t *Rectangle) {
	a.x = float64(t.B) / 2
	a.y = float64(t.L) / 2

	fmt.Printf("Calculating middle point coordinates for rectangle:%+v\n", a)
}
