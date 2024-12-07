package gotype

import (
	"github.com/benbjohnson/immutable"
)

type GoString struct {
	// TODO make color private
	Color     Color
	Stones    *immutable.Set[Point]
	Liberties *immutable.Set[Point]
}

func NewGoString(color Color, stones *immutable.Set[Point], liberties *immutable.Set[Point]) GoString {
	return GoString{
		Color:     color,
		Stones:    stones,
		Liberties: liberties,
	}
}

func (g *GoString) WithoutLiberty(point Point) *GoString {
	liberties := g.Liberties.Delete(point)
	return &GoString{
		Color:     g.Color,
		Stones:    g.Stones,
		Liberties: &liberties,
	}
}

func (g *GoString) WithLiberty(point Point) *GoString {
	liberties := g.Liberties.Add(point)
	return &GoString{
		Color:     g.Color,
		Stones:    g.Stones,
		Liberties: &liberties,
	}
}

func (g *GoString) MergeWith(other GoString) GoString {
	if g.Color != other.Color {
		panic("Cannot merge two strings of different colors")
	}

	stones := mergeSets(g.Stones, other.Stones)
	liberties := mergeSets(g.Liberties, other.Liberties)
	liberties = diffSets(liberties, stones)
	return GoString{
		Color:     g.Color,
		Stones:    stones,
		Liberties: liberties,
	}
}

func (g *GoString) NumLiberties() int {
	return g.Liberties.Len()
}

func (g *GoString) Equals(other GoString) bool {
	return g.Color == other.Color &&
		equalsSets(g.Stones, other.Stones) &&
		equalsSets(g.Liberties, other.Liberties)
}

func mergeSets[T comparable](set1, set2 *immutable.Set[T]) *immutable.Set[T] {
	mergedSet := *set1
	iter := set2.Iterator()
	for !iter.Done() {
		elem, _ := iter.Next()
		mergedSet = mergedSet.Add(elem)
	}
	return &mergedSet
}

func diffSets[T comparable](set1, set2 *immutable.Set[T]) *immutable.Set[T] {
	diffSet := *set1
	iter := set2.Iterator()
	for !iter.Done() {
		elem, _ := iter.Next()
		diffSet = diffSet.Delete(elem)
	}
	return &diffSet
}

// equalsSets checks if two immutable sets are equal
func equalsSets[T comparable](set1, set2 *immutable.Set[T]) bool {
	if set1.Len() != set2.Len() {
		return false
	}
	iter := set1.Iterator()
	for !iter.Done() {
		elem, _ := iter.Next()
		if !set2.Has(elem) {
			return false
		}
	}
	return true
}
