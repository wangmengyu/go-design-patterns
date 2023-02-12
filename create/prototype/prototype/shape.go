package prototype

type IShape interface {
	Clone() IShape
}

type Circle struct {
	X, Y, Radius int
}

func (c *Circle) Clone() IShape {
	return &Circle{
		X:      c.X,
		Y:      c.Y,
		Radius: c.Radius,
	}
}

type Rectangle struct {
	X1, Y1, X2, Y2 int
}

func (r *Rectangle) Clone() IShape {
	return &Rectangle{
		X1: r.X1,
		Y1: r.Y1,
		X2: r.X2,
		Y2: r.Y2,
	}
}
