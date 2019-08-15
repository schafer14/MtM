package board

import (
	"github.com/schafer14/MtM/common"
	"github.com/schafer14/MtM/move"
)

func (b Board) rookMoves(movesSlice *MoveList) {
	occ := b.Colors[0] | b.Colors[1]
	friendlies := b.Colors[b.Turn]
	allRooks := b.Pieces[common.Rook] & friendlies

	for rooks := allRooks; rooks != 0; rooks &= rooks - 1 {
		src := common.FirstOne(rooks)

		blocker := occ & rookMagic[src].mask
		index := (blocker * rookMagic[src].magic) >> 52
		moves := rookMagicMoves[src][index]

		allLegalMoves := moves & ^friendlies

		for legalMoves := allLegalMoves; legalMoves != 0; legalMoves &= legalMoves - 1 {
			dest := common.FirstOne(legalMoves)
			isCap, capPiece := b.pieceOn(dest)
			if isCap {
				movesSlice.Append(move.Move32(src | dest<<6 | common.Rook<<13 | capPiece<<16 | dest<<19 | 1<<25))
			} else {
				movesSlice.Append(move.Move32(src | dest<<6 | common.Rook<<13))
			}
		}

	}

}

func (b Board) rookAttacks(turn uint) (attackSpace uint64) {
	occ := b.Colors[0] | b.Colors[1]
	friendlies := b.Colors[turn]
	allRooks := b.Pieces[common.Rook] & friendlies

	for rooks := allRooks; rooks != 0; rooks &= rooks - 1 {
		squareNum := common.FirstOne(rooks)

		blocker := occ & rookMagic[squareNum].mask
		index := (blocker * rookMagic[squareNum].magic) >> 52
		attackSpace |= rookMagicMoves[squareNum][index]
	}

	return attackSpace
}
