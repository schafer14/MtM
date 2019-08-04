package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func (b Board) queenMoves(movesSlice *MoveList) {
	occ := b.Colors[0] | b.Colors[1]
	friendlies := b.Colors[b.Turn]
	allQueens := b.Pieces[common.Queen] & friendlies

	for Queens := allQueens; Queens != 0; Queens &= Queens - 1 {
		src := common.FirstOne(Queens)

		// Stright Moves
		blocker := occ & rookMagic[src].mask
		index := (blocker * rookMagic[src].magic) >> 52
		moves := rookMagicMoves[src][index]

		allLegalMoves := moves & ^friendlies

		for legalMoves := allLegalMoves; legalMoves != 0; legalMoves &= legalMoves - 1 {
			dest := common.FirstOne(legalMoves)
			isCap, capPiece := b.pieceOn(dest)
			if isCap {
				movesSlice.Append(move.Move32(src | dest<<6 | common.Queen<<13 | capPiece<<16 | dest<<19 | 1<<25))
			} else {
				movesSlice.Append(move.Move32(src | dest<<6 | common.Queen<<13))
			}
		}

		// Diagonal Moves
		blocker2 := occ & bishopMagic[src].mask
		index2 := (blocker2 * bishopMagic[src].magic) >> 55
		moves2 := bishopMagicMoves[src][index2]

		allLegalMoves2 := moves2 & ^friendlies

		for legalMoves := allLegalMoves2; legalMoves != 0; legalMoves &= legalMoves - 1 {
			dest := common.FirstOne(legalMoves)
			isCap, capPiece := b.pieceOn(dest)
			if isCap {
				movesSlice.Append(move.Move32(src | dest<<6 | common.Queen<<13 | capPiece<<16 | dest<<19 | 1<<25))
			} else {
				movesSlice.Append(move.Move32(src | dest<<6 | common.Queen<<13))
			}
		}

	}

}

func (b Board) queenAttacks(turn uint) (attackSpace uint64) {
	occ := b.Colors[0] | b.Colors[1]
	friendlies := b.Colors[turn]
	allQueens := b.Pieces[common.Queen] & friendlies

	for Queens := allQueens; Queens != 0; Queens &= Queens - 1 {
		squareNum := common.FirstOne(Queens)

		// Stright Moves
		blocker := occ & rookMagic[squareNum].mask
		index := (blocker * rookMagic[squareNum].magic) >> 52
		attackSpace |= rookMagicMoves[squareNum][index]

		// Diagonal Moves
		blocker2 := occ & bishopMagic[squareNum].mask
		index2 := (blocker2 * bishopMagic[squareNum].magic) >> 55
		attackSpace |= bishopMagicMoves[squareNum][index2]
	}

	return attackSpace
}
