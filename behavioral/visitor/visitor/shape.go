package visitor

type Shape interface {
	// 形状接口
	GetType() string  // 获得当前形状类型
	Accept(v Visitor) // 接收访问者
}
type Square struct { // 方形
	Side int
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

type Rectangle struct {
	L int
	B int
}

func (t *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(t)
}

func (t *Rectangle) GetType() string {
	return "rectangle"
}
