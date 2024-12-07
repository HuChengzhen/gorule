package gotype

import "github.com/benbjohnson/immutable"

type Board struct {
	NumRows uint8
	NumCols uint8
	grid    map[Point]*GoString
	hash    int64
}

func NewBoard(numRows, numCols uint8) Board {
	return Board{
		NumRows: numRows,
		NumCols: numCols,
		grid:    make(map[Point]*GoString),
		hash:    EMPTY_BOARD,
	}
}

func (b *Board) IsOnGrid(point Point) bool {
	return 1 <= point.row && point.row <= b.NumRows &&
		1 <= point.col && point.col <= b.NumCols
}

func (b *Board) PlaceStone(color Color, point Point) {
	if !b.IsOnGrid(point) {
		panic("point is not on grid")
	}

	if b.grid[point] != nil {
		panic("point already occupied")
	}

	adjacentSameColor := immutable.NewSet[Point](PointHasher{})
	adjacentOppositeColor := immutable.NewSet[Point](PointHasher{})
	liberties := immutable.NewSet[Point](PointHasher{})

	for _, neighbor := range point.neighbors() {
		if !b.IsOnGrid(neighbor) {
			continue
		}

		neighborString := b.grid[neighbor]
		if neighborString == nil {
			liberties = liberties.Add(neighbor)
		} else if neighborString.Color == color {
			adjacentSameColor = adjacentSameColor.Add(neighbor)
		} else {
			adjacentOppositeColor = adjacentOppositeColor.Add(neighbor)
		}
	}

	add := immutable.NewSet[Point](PointHasher{}).Add(point)
	newString := NewGoString(color, &add, &liberties)

	iterator := adjacentSameColor.Iterator()
	for !iterator.Done() {
		sameColorString, _ := iterator.Next()
		newString = newString.MergeWith(*b.grid[sameColorString])
	}

	iterator = newString.Stones.Iterator()
	for !iterator.Done() {
		newStringPoint, _ := iterator.Next()
		b.grid[newStringPoint] = &newString
	}

	hash, _ := hashCodeMap.Get(PlayMove{
		Point: point,
		Color: color,
	})

	b.hash ^= hash

	iterator = adjacentOppositeColor.Iterator()
	for !iterator.Done() {
		otherColorString, _ := iterator.Next()
		otherString := b.grid[otherColorString]
		replacement := otherString.WithoutLiberty(point)
		if replacement.NumLiberties() > 0 {
			b.ReplaceString(otherString.WithoutLiberty(point))
		} else {
			b.RemoveString(otherString)
		}
	}
}

func (b *Board) ReplaceString(newString *GoString) {
	iterator := newString.Stones.Iterator()
	for !iterator.Done() {
		point, _ := iterator.Next()
		b.grid[point] = newString
	}
}

func (b *Board) RemoveString(string *GoString) {
	iterator := string.Stones.Iterator()
	for !iterator.Done() {
		point, _ := iterator.Next()
		for _, neighbors := range point.neighbors() {
			neighborString := b.grid[neighbors]
			if neighborString == nil {
				continue
			}

			if neighborString != string {
				b.ReplaceString(neighborString.WithLiberty(point))
			}
		}

		delete(b.grid, point)

		hash, _ := hashCodeMap.Get(PlayMove{
			Point: point,
			Color: string.Color,
		})

		b.hash ^= hash
	}
}

func (b *Board) GetColor(point Point) Color {
	goString, ok := b.grid[point]
	if !ok {
		return Empty
	}
	return goString.Color
}

func (b *Board) GetGoString(point Point) *GoString {
	goString, ok := b.grid[point]
	if !ok {
		return nil
	}
	return goString
}

func (b *Board) GetHash() int64 {
	return b.hash
}

func (b *Board) Copy() *Board {
	newBoard := NewBoard(b.NumRows, b.NumCols)
	for point, goString := range b.grid {
		newBoard.grid[point] = goString
	}
	newBoard.hash = b.hash
	return &newBoard
}
