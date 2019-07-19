package board

import (
	"testing"

	"github.com/schafer14/chess/move"
)

type queenTest struct {
	fen      string
	numMoves int
}

var queenTests []queenTest = []queenTest{
	queenTest{
		fen:      "8/8/8/8/8/8/8/Q7 w KQkq - 0 1",
		numMoves: 21,
	},
	queenTest{
		fen:      "p7/8/8/8/8/8/8/Q7 w KQkq - 0 1",
		numMoves: 21,
	},
	queenTest{
		fen:      "p7/p7/8/8/8/8/8/Q7 w KQkq - 0 1",
		numMoves: 20,
	},
	queenTest{
		fen:      "p7/P7/8/8/8/8/8/Q7 w KQkq - 0 1",
		numMoves: 19,
	},
	queenTest{
		fen:      "p7/P7/8/8/8/8/8/Q1p5 w KQkq - 0 1",
		numMoves: 14,
	},
	queenTest{
		fen:      "p7/P7/8/8/8/8/8/Q1P5 w KQkq - 0 1",
		numMoves: 13,
	},
}

func TestQueenMoves(t *testing.T) {
	for _, tt := range queenTests {
		var moves []move.Move32
		board := FromFen(tt.fen)

		moveStream := board.queenMoves()

		for move := range moveStream {
			moves = append(moves, move)
		}

		if len(moves) != tt.numMoves {
			t.Errorf("%v expected %v queen moves but got %v", tt.fen, tt.numMoves, len(moves))
		}
	}
}

func BenchmarkQueen(b *testing.B) {
	board := FromFen("8/8/8/1Q6/1P6/8/8/8 w - - 0 1")
	for i := 0; i < b.N; i++ {
		board.queenMoves()
	}
}
