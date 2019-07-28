package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func (b Board) knightMoves(movesSlice *MoveList) {
	friendlies := b.colors[b.turn]
	allKnights := b.pieces[common.Knight] & friendlies
	knightMover := move.Mover(common.Knight)

	for knights := allKnights; knights != 0; knights &= knights - 1 {
		squareNum := common.FirstOne(knights)
		mover, capMover := knightMover(squareNum)

		moves := knightAttacks[squareNum]
		legalMoves := moves & ^friendlies

		for ; legalMoves != 0; legalMoves &= legalMoves - 1 {
			dest := common.FirstOne(legalMoves)
			isCap, capPiece := b.pieceOn(dest)
			if isCap {
				movesSlice.Append(capMover(dest, capPiece))
			} else {
				movesSlice.Append(mover(dest))
			}
		}

	}
}

func (b Board) knightAttacks(turn uint) (attackSpace uint64) {
	friendlies := b.colors[turn]
	allKnights := b.pieces[common.Knight] & friendlies

	for knights := allKnights; knights != 0; knights &= knights - 1 {
		squareNum := common.FirstOne(knights)

		attackSpace |= knightAttacks[squareNum]
	}

	return attackSpace
}
