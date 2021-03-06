package board

import (
	"github.com/schafer14/MtM/common"
	"github.com/schafer14/MtM/move"
)

func (b Board) bishopMoves(movesSlice *MoveList) {
	occ := b.Colors[0] | b.Colors[1]
	friendlies := b.Colors[b.Turn]
	allBishops := b.Pieces[common.Bishop] & friendlies

	for bishops := allBishops; bishops != 0; bishops &= bishops - 1 {
		src := common.FirstOne(bishops)

		blocker := occ & bishopMagic[src].mask
		index := (blocker * bishopMagic[src].magic) >> 55
		moves := bishopMagicMoves[src][index]

		allLegalMoves := moves & ^friendlies

		for legalMoves := allLegalMoves; legalMoves != 0; legalMoves &= legalMoves - 1 {
			dest := common.FirstOne(legalMoves)
			isCap, capPiece := b.pieceOn(dest)
			if isCap {
				movesSlice.Append(move.Move32(src | dest<<6 | common.Bishop<<13 | capPiece<<16 | dest<<19 | 1<<25))
			} else {
				movesSlice.Append(move.Move32(src | dest<<6 | common.Bishop<<13))
			}
		}

	}

}

func (b Board) bishopAttacks(turn uint) (attackSpace uint64) {
	occ := b.Colors[0] | b.Colors[1]
	friendlies := b.Colors[turn]
	allBishops := b.Pieces[common.Bishop] & friendlies

	for bishops := allBishops; bishops != 0; bishops &= bishops - 1 {
		squareNum := common.FirstOne(bishops)

		blocker := occ & bishopMagic[squareNum].mask
		index := (blocker * bishopMagic[squareNum].magic) >> 55
		attackSpace |= bishopMagicMoves[squareNum][index]
	}

	return attackSpace
}
