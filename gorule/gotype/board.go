package gotype

type Board struct {
	NumRows int8
	NumCols int8
	grid    map[Point]GoString
}

func (b *Board) PlaceStone(color Color, point Point) {

}
