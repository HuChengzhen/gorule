package gotype

import "fmt"

type Point struct {
	row, col uint8
}

func NewPoint(row, col uint8) Point {
	return Point{row, col}
}

func (p Point) neighbors() []Point {
	return []Point{
		{p.row - 1, p.col},
		{p.row + 1, p.col},
		{p.row, p.col - 1},
		{p.row, p.col + 1},
	}
}

func (p Point) String() string {
	return fmt.Sprintf("Point (%d, %d)", p.row, p.col)
}

type PointHasher struct{}

func (h PointHasher) Hash(p Point) uint32 {
	return uint32(p.row)<<16 | uint32(p.col)
}

func (h PointHasher) Equal(a, b Point) bool {
	return a == b
}
