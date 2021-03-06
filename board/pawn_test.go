package board

import (
	"testing"
)

type pawnTest struct {
	fen      string
	numMoves int
}

var pawnTests []pawnTest = []pawnTest{
	// White Pieces
	pawnTest{
		fen:      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		numMoves: 16,
	},
	pawnTest{
		fen:      "rnbqkbnr/pppppppp/8/8/p7/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		numMoves: 15,
	},
	pawnTest{
		fen:      "rnbqkbnr/pppppppp/8/8/pppppppp/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		numMoves: 8,
	},
	pawnTest{
		fen:      "8/8/8/pppppppp/P7/8/8/8 w KQkq - 0 1",
		numMoves: 1,
	},
	pawnTest{
		fen:      "8/8/8/pppppppp/1P6/8/8/8 w KQkq - 0 1",
		numMoves: 2,
	},
	pawnTest{
		fen:      "8/8/8/p1pppppp/1P6/8/8/8 w KQkq - 0 1",
		numMoves: 3,
	},
	pawnTest{
		fen:      "8/8/8/8/1P6/8/8/8 w KQkq - 0 1",
		numMoves: 1,
	},
	pawnTest{
		fen:      "8/8/8/pppppppp/7P/8/8/8 w KQkq - 0 1",
		numMoves: 1,
	},
	// Black Pieces
	pawnTest{
		fen:      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
		numMoves: 16,
	},
	pawnTest{
		fen:      "rnbqkbnr/pppppppp/8/P7/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
		numMoves: 15,
	},
	pawnTest{
		fen:      "rnbqkbnr/pppppppp/8/PPPPPPPP/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
		numMoves: 8,
	},
	pawnTest{
		fen:      "8/8/8/p7/PPPPPPPP/8/8/8 b KQkq - 0 1",
		numMoves: 1,
	},
	pawnTest{
		fen:      "8/8/8/1p6/PPPPPPPP/8/8/8 b KQkq - 0 1",
		numMoves: 2,
	},
	pawnTest{
		fen:      "8/8/8/1p6/P1PPPPPP/8/8/8 b KQkq - 0 1",
		numMoves: 3,
	},
	pawnTest{
		fen:      "8/8/8/8/1p6/8/8/8 b KQkq - 0 1",
		numMoves: 1,
	},
	pawnTest{
		fen:      "8/8/8/7p/PPPPPPPP/8/8/8 b KQkq - 0 1",
		numMoves: 1,
	},
	// EnPassant
	pawnTest{
		fen:      "8/8/8/8/3p4/3P4/8/8 b KQkq e3 0 1",
		numMoves: 1,
	},
	pawnTest{
		fen:      "8/8/3p4/3P4/8/8/8/8 w KQkq c6 0 1",
		numMoves: 1,
	},
}

func TestPawnMoves(t *testing.T) {
	for _, tt := range pawnTests {
		var moves MoveList
		board := FromFen(tt.fen)
		moves.Reset()

		board.pawnMoves(&moves)

		if moves.Len() != tt.numMoves {
			t.Errorf("%v expected %v pawn moves but got %v", tt.fen, tt.numMoves, moves.Len())
		}
	}
}
