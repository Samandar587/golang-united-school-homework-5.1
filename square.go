package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End(a, b int) Point {
	return Point{x: a - int(s.a), y: b - int(s.a)}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return 4 * s.a
}
