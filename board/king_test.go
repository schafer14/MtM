package board

import (
	"testing"
)

type kingTest struct {
	fen      string
	numMoves int
}

var kingTests []kingTest = []kingTest{
	kingTest{
		fen:      "8/8/8/8/8/8/8/1K6 w kq - 0 1",
		numMoves: 5,
	},
	kingTest{
		fen:      "8/8/8/3K4/8/8/8/8 w kq - 0 1",
		numMoves: 8,
	},
	kingTest{
		fen:      "8/8/3p4/3K4/3P4/8/8/8 w kq - 0 1",
		numMoves: 7,
	},
}

func TestKingMoves(t *testing.T) {
	for _, tt := range kingTests {
		var moves MoveList
		board := FromFen(tt.fen)

		moves.Reset()
		board.kingMoves(&moves)

		if moves.Len() != tt.numMoves {
			t.Errorf("%v expected %v king moves but got %v", tt.fen, tt.numMoves, moves.Len())
		}
	}
}
