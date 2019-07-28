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

		board.bishopMoves(&moves)

		if len(moves) != tt.numMoves {
			t.Errorf("%v expected %v bishop moves but got %v", tt.fen, tt.numMoves, len(moves))
		}
	}
}

func BenchmarkBishopMoves(b *testing.B) {
	var moves []move.Move32
	board := FromFen("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")
	for i := 0; i < b.N; i++ {
		board.bishopMoves(&moves)
	}
}
