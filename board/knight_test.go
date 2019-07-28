package board

import (
	"testing"
)

type knightTest struct {
	fen      string
	numMoves int
}

var knightTests []knightTest = []knightTest{
	knightTest{
		fen:      "8/8/8/8/8/8/8/1N6 w KQkq - 0 1",
		numMoves: 3,
	},
	knightTest{
		fen:      "8/8/8/8/8/p7/PPPP4/1N6 w KQkq - 0 1",
		numMoves: 2,
	},
	knightTest{
		fen:      "8/8/8/8/8/P7/PPPP4/1N6 w KQkq - 0 1",
		numMoves: 1,
	},
}

func TestKnightMoves(t *testing.T) {
	for _, tt := range knightTests {
		var moves MoveList
		board := FromFen(tt.fen)

		moves.Reset()

		board.knightMoves(&moves)

		if moves.Len() != tt.numMoves {
			t.Errorf("%v expected %v knight moves but got %v", tt.fen, tt.numMoves, moves.Len())
		}
	}
}
