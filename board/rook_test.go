package board

import (
	"testing"
)

type rookTest struct {
	fen      string
	numMoves int
}

var rookTests []rookTest = []rookTest{
	rookTest{
		fen:      "8/8/8/8/8/8/8/R7 w KQkq - 0 1",
		numMoves: 14,
	},
	rookTest{
		fen:      "p7/8/8/8/8/8/8/R7 w KQkq - 0 1",
		numMoves: 14,
	},
	rookTest{
		fen:      "p7/p7/8/8/8/8/8/R7 w KQkq - 0 1",
		numMoves: 13,
	},
	rookTest{
		fen:      "p7/P7/8/8/8/8/8/R7 w KQkq - 0 1",
		numMoves: 12,
	},
	rookTest{
		fen:      "p7/P7/8/8/8/8/8/R1p5 w KQkq - 0 1",
		numMoves: 7,
	},
	rookTest{
		fen:      "p7/P7/8/8/8/8/8/R1P5 w KQkq - 0 1",
		numMoves: 6,
	},
}

func TestRookMoves(t *testing.T) {
	for _, tt := range rookTests {
		var moves MoveList
		board := FromFen(tt.fen)

		moves.Reset()
		board.rookMoves(&moves)

		if moves.Len() != tt.numMoves {
			t.Errorf("%v expected %v rook moves but got %v", tt.fen, tt.numMoves, moves.Len())
		}
	}
}
