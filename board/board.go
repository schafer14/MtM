package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

// Moves generates a list of legal moves in a position
func (b Board) Moves() []move.Move32 {
	// TODO optimise this when there is interent
	moves := make([]move.Move32, 0, 256)

	for _, move := range b.PsudoMoves() {
		// Optimsie this when the undo move func is done
		test := b.Clone()
		test.Move(move)
		king := test.colors[test.opp()] & test.pieces[common.King]
		oppAttack := test.attackSpace(test.turn)

		if oppAttack&king == 0 {
			moves = append(moves, move)
		}
	}

	return moves
}

// Generates a list of moves without checking for moving into check
func (b Board) PsudoMoves() []move.Move32 {
	moves := make([]move.Move32, 0, 256)

	b.pawnMoves(&moves)
	b.knightMoves(&moves)
	b.bishopMoves(&moves)
	b.rookMoves(&moves)
	b.queenMoves(&moves)
	b.kingMoves(&moves)

	return moves
}

// Clone duplicates a chess board.
func (b Board) Clone() Board {
	newBoard := b
	return newBoard
}

// New creates a new board initialized to the intial position.
func New() Board {
	return initialBoard.Clone()
}

// Empty creates a board with no pieces on it.
func Empty() Board {
	return Board{}
}

// FromFen creates a new board from a fenstring.
var FromFen = fromFen

func (b *Board) applyMoves(moves []string) {
	for _, moveStr := range moves {
		move, _ := b.MoveFromSrcDestNotation(moveStr)
		b.Move(move)
	}
}

func (b Board) attackSpace(turn uint) uint64 {
	return b.pawnAttacks(turn) |
		b.knightAttacks(turn) |
		b.bishopAttacks(turn) |
		b.rookAttacks(turn) |
		b.queenAttacks(turn) |
		b.kingAttacks(turn)
}
