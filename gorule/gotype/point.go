package gotype

import "fmt"

type Point struct {
	row, col uint8
}

func (p Point) neighbors() []Point {
	return []Point{
		Point{p.row - 1, p.col},
		Point{p.row + 1, p.col},
		Point{p.row, p.col - 1},
		Point{p.row, p.col + 1},
	}
}

func (p Point) String() string {
	return fmt.Sprintf("Point (%d, %d)", p.row, p.col)
}
