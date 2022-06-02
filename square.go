package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (p Point) End(X, Y int) Point {
	return Point{x: X, y: Y}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return 4 * s.a
}
