package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func (b Board) knightMoves(movesSlice *[]move.Move32) {
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
				*movesSlice = append(*movesSlice, capMover(dest, capPiece))
			} else {
				*movesSlice = append(*movesSlice, mover(dest))
			}
		}

	}
}
