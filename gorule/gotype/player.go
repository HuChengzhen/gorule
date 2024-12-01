package gotype

import "log"

type Color int8

const (
	_ = iota // iota generates incrementing numbers starting from 0
	Black
	White
)

func (c Color) Other() Color {
	if c == Black {
		return White
	} else if c == White {
		return Black
	}
	log.Panicf("Color is invalid %d", c)
	return 0
}

func (c Color) String() string {
	if c == Black {
		return "Black"
	} else if c == White {
		return "White"
	}
	log.Panicf("Color is invalid %d", c)
	return ""
}
