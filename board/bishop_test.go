package board

import (
	"testing"

	"github.com/schafer14/chess/move"
)

type bishopTest struct {
	fen      string
	numMoves int
}

var bishopTests []bishopTest = []bishopTest{
	bishopTest{
		fen:      "8/8/8/8/8/8/8/B7 w KQkq - 0 1",
		numMoves: 7,
	},
	bishopTest{
		fen:      "8/8/8/4B3/8/8/8/8 w KQkq - 0 1",
		numMoves: 13,
	},
	bishopTest{
		fen:      "1P7/8/8/4B3/8/8/8/8 w KQkq - 0 1",
		numMoves: 12,
	},
	bishopTest{
		fen:      "1p7/8/8/4B3/8/8/8/8 w KQkq - 0 1",
		numMoves: 13,
	},
	bishopTest{
		fen:      "1p7/8/3P4/4B3/8/8/8/8 w KQkq - 0 1",
		numMoves: 10,
	},
}

func TestBishopMoves(t *testing.T) {
	for _, tt := range bishopTests {
		var moves []move.Move32
		board := FromFen(tt.fen)

		moveStream := board.bishopMoves()

		for move := range moveStream {
			moves = append(moves, move)
		}

		if len(moves) != tt.numMoves {
			t.Errorf("%v expected %v bishop moves but got %v", tt.fen, tt.numMoves, len(moves))
		}
	}
}

func BenchmarkBishop(b *testing.B) {
	board := New()
	for i := 0; i < b.N; i++ {
		board.bishopMoves()
	}
}
