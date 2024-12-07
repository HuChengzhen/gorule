package gotype

import "github.com/benbjohnson/immutable"

type GameState struct {
	Board            *Board
	Next             Color
	PreviousState    *GameState
	PreviousStateSet *immutable.Set[HashState]
	LastMove         *Move
}

type HashState struct {
	NextColor Color
	Hash      int64
}

type HashStateHasher struct{}

func (h HashStateHasher) Hash(p HashState) uint32 {
	return uint32(p.NextColor)<<16 | uint32(p.Hash)
}

func (h HashStateHasher) Equal(a, b HashState) bool {
	return a == b
}

func NewGameState(board *Board, next Color, previousState *GameState, lastMove *Move) *GameState {
	var previousStateSet immutable.Set[HashState]

	if previousState == nil {
		previousStateSet = immutable.NewSet[HashState](HashStateHasher{})
	} else {
		previousStateSet = previousState.PreviousStateSet.Add(HashState{
			NextColor: previousState.Next,
			Hash:      previousState.Board.GetHash(),
		})
	}
	return &GameState{
		Board:            board,
		Next:             next,
		PreviousState:    previousState,
		PreviousStateSet: &previousStateSet,
		LastMove:         lastMove,
	}
}

func (g *GameState) ApplyMove(move *Move) *GameState {
	var nextBoard *Board

	if move.IsPlay() {
		nextBoard = g.Board.Copy()
		nextBoard.PlaceStone(g.Next, move.Point)
	} else {
		nextBoard = g.Board
	}
	return NewGameState(nextBoard, g.Next.Other(), g, move)
}

func NewGame(boardSize uint8) *GameState {
	board := NewBoard(boardSize, boardSize)
	return NewGameState(&board, Black, nil, nil)
}

func (g *GameState) IsValidMove(move Move) bool {
	if g.isOver() {
		return false
	}

	if move.IsPass || move.IsResign {
		return true
	}

	return g.Board.GetColor(move.Point) == Empty && !g.IsMoveSelfCapture(g.Next, move) && !g.IsMoveViolateKo(g.Next, move)

}

func (g *GameState) isOver() bool {
	if g.LastMove == nil {
		return false
	}

	if g.LastMove.IsResign {
		return true
	}

	move := g.PreviousState.LastMove
	if move == nil {
		return false
	}

	return g.LastMove.IsPass && move.IsPass
}

func (g *GameState) IsMoveSelfCapture(next Color, move Move) bool {
	if !move.IsPlay() {
		return false
	}

	nextBoard := g.Board.Copy()
	nextBoard.PlaceStone(next, move.Point)
	newString := nextBoard.GetGoString(move.Point)
	return newString.NumLiberties() == 0
}

func (g *GameState) IsMoveViolateKo(color Color, move Move) bool {
	if !move.IsPlay() {
		return false
	}

	nextBoard := g.Board.Copy()
	nextBoard.PlaceStone(color, move.Point)
	nextHash := nextBoard.GetHash()
	return g.PreviousStateSet.Has(HashState{
		NextColor: color.Other(),
		Hash:      nextHash,
	})
}
