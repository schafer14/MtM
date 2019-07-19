package board

import (
	"testing"

	"github.com/schafer14/chess/common"
)

func TestMoves(t *testing.T) {
	board := New()

	moves := board.Moves()

	if len(moves) != 20 {
		t.Errorf("Should be 20 moves available for first move got %v", len(moves))
	}
}

func TestEmpty(t *testing.T) {
	board := Empty()
	if board.colors[common.White] != 0 {
		t.Error("Pawns should be set to 0 when a empty board is created")
	}
	if board.colors[common.Black] != 0 {
		t.Error("Pawns should be set to 0 when a empty board is created")
	}
}

func BenchmarkEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Empty()
	}
}

func TestNew(t *testing.T) {
	board := New()
	if board.colors[common.White] != 0xffff {
		t.Error("white pieces should be initialized correclty when creating a new board")
	}

	if board.colors[common.Black] != 0xffff000000000000 {
		t.Error("black pieces should be initialized correclty when creating a new board")
	}

	if board.pieces[common.Pawn] != 0x00ff00000000ff00 {
		t.Error("pawns should be initialized correclty when creating a new board")
	}

	if board.pieces[common.Knight] != 0x4200000000000042 {
		t.Error("knights should be initialized correclty when creating a new board")
	}

	if board.pieces[common.Bishop] != 0x2400000000000024 {
		t.Error("bishops should be initialized correclty when creating a new board")
	}

	if board.pieces[common.Rook] != 0x8100000000000081 {
		t.Error("rooks should be initialized correclty when creating a new board")
	}

	if board.pieces[common.Queen] != 0x0800000000000008 {
		t.Error("queens should be initialized correclty when creating a new board")
	}

	if board.pieces[common.King] != 0x1000000000000010 {
		t.Error("kings should be initialized correclty when creating a new board")
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func TestClone(t *testing.T) {
	board := Board{colors: [2]uint64{0123, 43}}

	newBoard := board.Clone()

	if newBoard.colors[common.White] != 0123 {
		t.Error("clone should correctly clone a board")
	}
}

func BenchmarkClone(b *testing.B) {
	board := New()

	for i := 0; i < b.N; i++ {
		_ = board.Clone()
	}
}

type fenTest struct {
	fen   string
	board Board
}

func TestFromFen(t *testing.T) {
	fenTests := []fenTest{
		fenTest{
			fen:   "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			board: New(),
		},
		fenTest{
			fen: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
			board: Board{
				colors: [2]uint64{
					0x1000EFFF,
					0xFFFF000000000000,
				},
				pieces: [6]uint64{
					0x00FF00001000EF00,
					0x4200000000000042,
					0x2400000000000024,
					0x8100000000000081,
					0x0800000000000008,
					0x1000000000000010,
				},
				turn:      common.Black,
				castling:  [4]bool{true, true, true, true},
				enPassant: 21,
			},
		},
		fenTest{
			fen: "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b Qkq - 1 2",
			board: Board{
				colors: [2]uint64{
					0x1020EFBF,
					0xFFFB000400000000,
				},
				pieces: [6]uint64{
					0x00FB00041000EF00,
					0x4200000000200002,
					0x2400000000000024,
					0x8100000000000081,
					0x0800000000000008,
					0x1000000000000010,
				},
				turn:      common.Black,
				castling:  [4]bool{false, true, true, true},
				enPassant: 0,
			},
		},
		fenTest{
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1",
			board: Board{
				turn:      common.White,
				castling:  [4]bool{false, false, false, false},
				enPassant: 0,
				pieces: [6]uint64{
					0x00FF00000000FF00,
					0x4200000000000042,
					0x2400000000000024,
					0x8100000000000081,
					0x0800000000000008,
					0x1000000000000010,
				},
				colors: [2]uint64{
					0xFFFF,
					0xFFFF000000000000,
				},
			},
		},
	}

	for _, tt := range fenTests {
		board := FromFen(tt.fen)
		if board != tt.board {
			t.Errorf("Error parsing fen '%v'", tt.fen)
		}
	}
}
