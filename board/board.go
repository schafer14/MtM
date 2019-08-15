package board

import (
	"github.com/schafer14/MtM/common"
)

// Moves generates a list of legal moves in a position

func (b Board) Moves() (moves MoveList) {
	// TODO optimise this when there is interent
	var ml MoveList
	var test Board
	b.PsudoMoves(&ml)

	for {
		// Optimsie this when the undo move func is done
		hasNext, move := ml.Next()
		if !hasNext {
			break
		}

		test = b
		test.Move(move)

		if !test.IsInCheck(test.opp()) {
			moves.Append(move)
		}
	}

	return moves
}

func (b Board) IsInCheck(turn uint) bool {
	king := b.Colors[turn] & b.Pieces[common.King]
	opp := common.Black
	if turn == common.Black {
		opp = common.White
	}
	oppAttack := b.attackSpace(opp)

	return oppAttack&king != 0
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

// New creates a new board initialized to the intial position.
func New() Board {
	return initialBoard
}

// Empty creates a board with no pieces on it.
func Empty() Board {
	return Board{}
}

// FromFen creates a new board from a fenstring.
var FromFen = fromFen

func (b *Board) ApplyMoves(moves []string) {
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
