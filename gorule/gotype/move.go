package gotype

import "fmt"

type Move struct {
	Point    Point
	IsPass   bool
	IsResign bool
}

func (m *Move) IsPlay() bool {
	return m.Point != Point{}
}

func (m *Move) String() string {
	if m.IsPass {
		return "pass"
	}
	if m.IsResign {
		return "resign"
	}
	return fmt.Sprintf("Point (%d, %d)", m.Point.row, m.Point.col)
}

func NewPlayMove(point Point) Move {
	return Move{
		Point: point,
	}
}

func NewPassTurnMove() Move {
	return Move{
		IsPass: true,
	}
}

func NewResignMove() Move {
	return Move{
		IsResign: true,
	}
}
