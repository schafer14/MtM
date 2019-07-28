package board

import (
	"testing"
)

func TestMoveFromSrcDestNotation(t *testing.T) {
	errorTestMoveSrcDestNotationTests(t)
	successTestMoveSrcDestNotationTests(t)
}

type testErrorMoveSrcDestNotationShape struct {
	fen  string
	move string
}

func successTestMoveSrcDestNotationTests(t *testing.T) {
	tests := []testErrorMoveSrcDestNotationShape{
		testErrorMoveSrcDestNotationShape{
			fen:  "startpos",
			move: "e2e4",
		},
	}

	for _, tt := range tests {
		b := FromFen(tt.fen)

		_, err := b.MoveFromSrcDestNotation(tt.move)

		if err != nil {
			t.Errorf("Reading move '%v' from position '%v' should not produce an error but produced %v", tt.move, tt.fen, err)
		}
	}
}

func errorTestMoveSrcDestNotationTests(t *testing.T) {
	tests := []testErrorMoveSrcDestNotationShape{
		testErrorMoveSrcDestNotationShape{
			fen:  "startpos",
			move: "e2e4Q!",
		},
		testErrorMoveSrcDestNotationShape{
			fen:  "startpos",
			move: "x2e4",
		},
		testErrorMoveSrcDestNotationShape{
			fen:  "startpos",
			move: "e0e4",
		},
		testErrorMoveSrcDestNotationShape{
			fen:  "startpos",
			move: "e2o4",
		},
		testErrorMoveSrcDestNotationShape{
			fen:  "startpos",
			move: "e2e9",
		},
		testErrorMoveSrcDestNotationShape{
			fen:  "startpos",
			move: "e2e5",
		},
	}

	for _, tt := range tests {
		b := FromFen(tt.fen)

		_, err := b.MoveFromSrcDestNotation(tt.move)

		if err == nil {
			t.Errorf("Reading move '%v' from position '%v' should produce an error", tt.move, tt.fen)
		}
	}
}
