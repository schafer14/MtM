package board

import (
	"testing"

	"github.com/schafer14/chess/move"
)

type kingTest struct {
	fen      string
	numMoves int
}

var kingTests []kingTest = []kingTest{
	kingTest{
		fen:      "8/8/8/8/8/8/8/1K6 w KQkq - 0 1",
		numMoves: 5,
	},
	kingTest{
		fen:      "8/8/8/3K4/8/8/8/8 w KQkq - 0 1",
		numMoves: 8,
	},
	kingTest{
		fen:      "8/8/3p4/3K4/3P4/8/8/8 w KQkq - 0 1",
		numMoves: 7,
	},
}

func TestKingMoves(t *testing.T) {
	for _, tt := range kingTests {
		var moves []move.Move32
		board := FromFen(tt.fen)

		moveStream := board.kingMoves()

		for move := range moveStream {
			moves = append(moves, move)
		}

		if len(moves) != tt.numMoves {
			t.Errorf("%v expected %v king moves but got %v", tt.fen, tt.numMoves, len(moves))
		}
	}
}

func BenchmarkKing(b *testing.B) {
	board := New()
	for i := 0; i < b.N; i++ {
		board.kingMoves()
	}
}
