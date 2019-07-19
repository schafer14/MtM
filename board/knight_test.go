package board

import (
	"testing"

	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
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
		var moves []move.Move32
		board := FromFen(tt.fen)

		moveStream := board.knightMoves()

		for move := range moveStream {
			moves = append(moves, move)
		}

		if len(moves) != tt.numMoves {
			t.Errorf("%v expected %v knight moves but got %v", tt.fen, tt.numMoves, len(moves))
		}
	}
}

func BenchmarkFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		common.FirstOne(0xB000000000000000)
	}
}

func BenchmarkKnight(b *testing.B) {
	board := New()
	for i := 0; i < b.N; i++ {
		board.knightMoves()
	}
}
