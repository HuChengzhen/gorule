package gotype

import (
	mapset "github.com/deckarep/golang-set/v2"
	"log"
)

type GoString struct {
	Color     Color
	Stones    mapset.Set[Point]
	Liberties mapset.Set[Point]
}

func (s *GoString) RemoveLiberty(point Point) {
	s.Liberties.Remove(point)
}

func (s *GoString) AddLiberty(point Point) {
	s.Liberties.Add(point)
}

func (s *GoString) Equal(other *GoString) bool {
	// Compare Color
	if s.Color != other.Color {
		return false
	}

	// Compare Stones
	if !s.Stones.Equal(other.Stones) {
		return false
	}

	// Compare Liberties
	if !s.Liberties.Equal(other.Liberties) {
		return false
	}

	return true
}

func (s *GoString) MergedWith(o *GoString) GoString {
	if s.Color != o.Color {
		log.Panic("Wrong Go String color")
		return GoString{}
	}

	combinedStone := s.Stones.Union(o.Stones)

	return GoString{
		Color:     s.Color,
		Stones:    combinedStone,
		Liberties: s.Liberties.Union(s.Liberties).Difference(combinedStone),
	}
}

func (s *GoString) NumLiberties() int {
	return s.Liberties.Cardinality()
}
