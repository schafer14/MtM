package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func (b Board) queenMoves() <-chan move.Move32 {
	moveStream := make(chan move.Move32)
	occ := b.colors[0] | b.colors[1]
	friendlies := b.colors[b.turn]
	allQueens := b.pieces[common.Queen] & friendlies
	queenMover := move.Mover(common.Queen)

	go func() {
		defer close(moveStream)

		for Queens := allQueens; Queens != 0; Queens &= Queens - 1 {
			squareNum := common.FirstOne(Queens)
			mover, capMover := queenMover(squareNum)

			// Stright Moves
			blocker := occ & rookMagic[squareNum].mask
			index := (blocker * rookMagic[squareNum].magic) >> 52
			moves := rookMagicMoves[squareNum][index]

			allLegalMoves := moves & ^friendlies

			for legalMoves := allLegalMoves; legalMoves != 0; legalMoves &= legalMoves - 1 {
				dest := common.FirstOne(legalMoves)
				isCap, capPiece := b.pieceOn(dest)
				if isCap {
					moveStream <- capMover(dest, capPiece)
				} else {
					moveStream <- mover(dest)
				}
			}

			// Diagonal Moves
			blocker2 := occ & bishopMagic[squareNum].mask
			index2 := (blocker2 * bishopMagic[squareNum].magic) >> 55
			moves2 := bishopMagicMoves[squareNum][index2]

			allLegalMoves2 := moves2 & ^friendlies

			for legalMoves := allLegalMoves2; legalMoves != 0; legalMoves &= legalMoves - 1 {
				dest := common.FirstOne(legalMoves)
				isCap, capPiece := b.pieceOn(dest)
				if isCap {
					moveStream <- capMover(dest, capPiece)
				} else {
					moveStream <- mover(dest)
				}
			}

		}
	}()

	return moveStream
}
