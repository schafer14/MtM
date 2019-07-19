package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func (b Board) bishopMoves() <-chan move.Move32 {
	moveStream := make(chan move.Move32)
	occ := b.colors[0] | b.colors[1]
	friendlies := b.colors[b.turn]
	allBishops := b.pieces[common.Bishop] & friendlies
	bishopMover := move.Mover(common.Bishop)

	go func() {
		defer close(moveStream)

		for bishops := allBishops; bishops != 0; bishops &= bishops - 1 {
			squareNum := common.FirstOne(bishops)
			mover, capMover := bishopMover(squareNum)

			blocker := occ & bishopMagic[squareNum].mask
			index := (blocker * bishopMagic[squareNum].magic) >> 55
			moves := bishopMagicMoves[squareNum][index]

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

		}
	}()

	return moveStream
}
