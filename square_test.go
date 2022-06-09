package square

import (
	"testing"
)

func TestArea(t *testing.T) {
	square := Square{
		a: 2,
	}
	expected := uint(4)
	actual := square.Area()

	if actual != expected {
		t.Errorf("Error in area, expected:%v, but was:%v", expected, actual)
	}
}

func TestPerimeter(t *testing.T) {
	square := Square{
		a: 2,
	}
	expected := uint(8)
	actual := square.Perimeter()

	if actual != expected {
		t.Errorf("Error in perimeter, expected:%v, but was:%v", expected, actual)
	}
}

func TestEnd(t *testing.T) {
	expected := Point{
		x: 3,
		y: -1,
	}

	square := Square{
		a: 2,
		start: Point{
			x: 1,
			y: 1,
		},
	}

	actual := square.End()

	if expected.x != actual.x {
		t.Errorf("x coordinate is incorrect, expected:%v but was:%v", expected.x, actual.x)
	}

	if expected.y != actual.y {
		t.Errorf("y coordinate is incorrect, expected:%v but was:%v", expected.y, actual.y)
	}
}
