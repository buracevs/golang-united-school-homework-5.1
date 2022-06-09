package square

import "math"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	result := Point{
		x: s.start.x + int(s.a),
		y: s.start.y + int(s.a),
	}
	return result
}

func (s Square) Area() uint {
	result := math.Pow(float64(s.a), 2)
	return uint(result)
}

func (s Square) Perimeter() uint {
	return uint(s.a * 4)
}
