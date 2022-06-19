// Package factory represents the use of the Factory Method pattern.
// The usecase for this is somewhat made up since the unexported fields of the
// shapes (circle, square and rectangle) can't really be updated/added when
// using this in Go in the way that I have created them. I personally tend to
// not use this pattern much in Go.
package factory

import (
	"fmt"
	"math"
)

// Shape the interface for all shapes.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type shape int

const (
	Circle shape = iota
	Square
	Rectangle
)

type rectangle struct {
	base   float64
	height float64
}

func (r *rectangle) Area() float64 {
	return r.base * r.height
}

func (r *rectangle) Perimeter() float64 {
	return 2 * (r.base + r.height)
}

type square struct {
	side float64
}

func (s *square) Area() float64 {
	return s.side * s.side
}

func (s *square) Perimeter() float64 {
	return 4 * s.side
}

type circle struct {
	rad float64
}

func (c *circle) Area() float64 {
	return math.Pi * c.rad * c.rad
}

func (c *circle) Perimeter() float64 {
	return 2 * math.Pi * c.rad
}

func ConstructShape(s shape) (Shape, error) {
	switch s {
	case Circle:
		return &circle{}, nil
	case Square:
		return &square{}, nil
	case Rectangle:
		return &rectangle{}, nil
	default:
		return nil, fmt.Errorf("unknown shape %v provided", s)
	}
}
