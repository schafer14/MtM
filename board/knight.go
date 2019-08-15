package board

import (
	"github.com/schafer14/MtM/common"
	"github.com/schafer14/MtM/move"
)

func (b Board) knightMoves(movesSlice *MoveList) {
	friendlies := b.Colors[b.Turn]
	allKnights := b.Pieces[common.Knight] & friendlies

	for knights := allKnights; knights != 0; knights &= knights - 1 {
		src := common.FirstOne(knights)

		moves := knightAttacks[src]
		legalMoves := moves & ^friendlies

		for ; legalMoves != 0; legalMoves &= legalMoves - 1 {
			dest := common.FirstOne(legalMoves)
			isCap, capPiece := b.pieceOn(dest)
			if isCap {
				movesSlice.Append(move.Move32(src | dest<<6 | common.Knight<<13 | capPiece<<16 | dest<<19 | 1<<25))
			} else {
				movesSlice.Append(move.Move32(src | dest<<6 | common.Knight<<13))
			}
		}

	}
}

func (b Board) knightAttacks(turn uint) (attackSpace uint64) {
	friendlies := b.Colors[turn]
	allKnights := b.Pieces[common.Knight] & friendlies

	for knights := allKnights; knights != 0; knights &= knights - 1 {
		squareNum := common.FirstOne(knights)

		attackSpace |= knightAttacks[squareNum]
	}

	return attackSpace
}
