package board

import (
	"testing"

	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func TestMove(t *testing.T) {
	board := New()

	// 1. d4
	board.Move(move.PawnQuiet(11, 27))

	if board.turn != common.Black {
		t.Errorf("Move should update color")
	}

	if board.colors[common.White]&board.pieces[common.Pawn] != 0x0800F700 {
		t.Errorf("After 1. d4 pawns should be 0x0800F700 got %#016x", board.colors[common.White]&board.pieces[common.Pawn])
	}

	if board.enPassant != 19 {
		t.Errorf("Move should add en Passant to board")
	}
}
