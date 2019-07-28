package board

import (
	"github.com/schafer14/chess/common"
)

// Moves generates a list of legal moves in a position
func (b Board) Moves() (moves MoveList) {
	// TODO optimise this when there is interent
	var ml MoveList
	b.PsudoMoves(&ml)

	for {
		// Optimsie this when the undo move func is done
		hasNext, move := ml.Next()
		if !hasNext {
			break
		}

		test := b.Clone()
		test.Move(move)
		king := test.colors[test.opp()] & test.pieces[common.King]
		oppAttack := test.attackSpace(test.turn)

		if oppAttack&king == 0 {
			moves.Append(move)
		}
	}

	return moves
}

// Generates a list of moves without checking for moving into check
func (b Board) PsudoMoves(ml *MoveList) {
	b.pawnMoves(ml)
	b.knightMoves(ml)
	b.bishopMoves(ml)
	b.rookMoves(ml)
	b.queenMoves(ml)
	b.kingMoves(ml)
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
