package square

type Point struct {
	x, y uint
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End(p Point) Point {
	return Point{x: p.x + s.a, y: p.x + s.a}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return 4 * s.a
}
