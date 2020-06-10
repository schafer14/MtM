package board

import (
	"testing"

	"github.com/schafer14/MtM/common"
)

func TestPromo(t *testing.T) {
	b := FromFen("8/P7/8/k7/8/K7/8/8 w KQkq - 0 1")
	if b.Moves().Len() != 7 {
		t.Error("Promotion is not working")
	}
}

func TestString(t *testing.T) {
	board := New()
	if board.String() != "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1" {
		t.Errorf("Board.String() failing to create fen string got %v", board.String())
	}

	fen := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"
	board = FromFen(fen)
	if board.String() != fen {
		t.Errorf("Fen %v failed to string got %v", fen, board.String())
	}

	fen = "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq h3 0 1"
	board = FromFen(fen)
	if board.String() != fen {
		t.Errorf("Fen %v failed to string got %v", fen, board.String())
	}

}

func TestMoves(t *testing.T) {
	board := New()

	moves := board.Moves()

	if moves.Len() != 20 {
		t.Errorf("Should be 20 moves available for first move got %v", moves.Len())
	}
}

func TestEmpty(t *testing.T) {
	board := Empty()
	if board.Colors[common.White] != 0 {
		t.Error("Pawns should be set to 0 when a empty board is created")
	}
	if board.Colors[common.Black] != 0 {
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
	if board.Colors[common.White] != 0xffff {
		t.Error("white pieces should be initialized correclty when creating a new board")
	}

	if board.Colors[common.Black] != 0xffff000000000000 {
		t.Error("black pieces should be initialized correclty when creating a new board")
	}

	if board.Pieces[common.Pawn] != 0x00ff00000000ff00 {
		t.Error("pawns should be initialized correclty when creating a new board")
	}

	if board.Pieces[common.Knight] != 0x4200000000000042 {
		t.Error("knights should be initialized correclty when creating a new board")
	}

	if board.Pieces[common.Bishop] != 0x2400000000000024 {
		t.Error("bishops should be initialized correclty when creating a new board")
	}

	if board.Pieces[common.Rook] != 0x8100000000000081 {
		t.Error("rooks should be initialized correclty when creating a new board")
	}

	if board.Pieces[common.Queen] != 0x0800000000000008 {
		t.Error("queens should be initialized correclty when creating a new board")
	}

	if board.Pieces[common.King] != 0x1000000000000010 {
		t.Error("kings should be initialized correclty when creating a new board")
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
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
				Colors: [2]uint64{
					0x1000EFFF,
					0xFFFF000000000000,
				},
				Pieces: [6]uint64{
					0x00FF00001000EF00,
					0x4200000000000042,
					0x2400000000000024,
					0x8100000000000081,
					0x0800000000000008,
					0x1000000000000010,
				},
				Turn:      common.Black,
				castling:  [4]bool{true, true, true, true},
				enPassant: 20,
			},
		},
		fenTest{
			fen: "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b Qkq - 1 2",
			board: Board{
				Colors: [2]uint64{
					0x1020EFBF,
					0xFFFB000400000000,
				},
				Pieces: [6]uint64{
					0x00FB00041000EF00,
					0x4200000000200002,
					0x2400000000000024,
					0x8100000000000081,
					0x0800000000000008,
					0x1000000000000010,
				},
				Turn:      common.Black,
				castling:  [4]bool{false, true, true, true},
				enPassant: 0,
			},
		},
		fenTest{
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1",
			board: Board{
				Turn:      common.White,
				castling:  [4]bool{false, false, false, false},
				enPassant: 0,
				Pieces: [6]uint64{
					0x00FF00000000FF00,
					0x4200000000000042,
					0x2400000000000024,
					0x8100000000000081,
					0x0800000000000008,
					0x1000000000000010,
				},
				Colors: [2]uint64{
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
