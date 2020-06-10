package move

import (
	"testing"

	"github.com/schafer14/MtM/common"
)

func TestGenerate(t *testing.T) {
	pawnMover := Mover(common.Pawn)
	e3Mover, capMove := pawnMover(23)

	straightMove := e3Mover(32)

	if straightMove.Src() != 23 {
		t.Errorf("Move generator expected src 23 got %v", straightMove.Src())
	}

	if straightMove.Dest() != 32 {
		t.Errorf("Move generator expected src 32 got %v", straightMove.Dest())
	}

	if straightMove.Piece() != common.Pawn {
		t.Errorf("Move generator expected src %v got %v", common.Pawn, straightMove.Piece())
	}

	if straightMove.IsCap() {
		t.Errorf("Capture bit should not be set")
	}

	capRight := capMove(33, common.Knight)

	if capRight.Src() != 23 {
		t.Errorf("Move generator expected src 23 got %v", capRight.Src())
	}

	if capRight.Dest() != 33 {
		t.Errorf("Move generator expected src 33 got %v", capRight.Dest())
	}

	if capRight.Piece() != common.Pawn {
		t.Errorf("Move generator expected src %v got %v", common.Pawn, capRight.Piece())
	}

	if !capRight.IsCap() {
		t.Errorf("Capture bit should be set")
	}

	capPiece, capSquare := capRight.Capture()

	if capPiece != common.Knight || capSquare != 33 {
		t.Errorf("Move generator expected src (1, 33) got (%v, %v)", capPiece, capSquare)
	}
}
