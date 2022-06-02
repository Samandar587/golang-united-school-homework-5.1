package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End(p Point) Point {
	return Point{x: int(s.a) + p.x, y: int(s.a) + p.y}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return 4 * s.a
}
