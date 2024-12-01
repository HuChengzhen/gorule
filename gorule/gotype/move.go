package gotype

type Move struct {
	Point    *Point
	IsPass   bool
	IsResign bool
}

func (m Move) IsPlay() bool {
	return m.Point != nil
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
