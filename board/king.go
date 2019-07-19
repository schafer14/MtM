package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func (b Board) kingMoves() <-chan move.Move32 {
	moveStream := make(chan move.Move32)
	friendlies := b.colors[b.turn]
	allKings := b.pieces[common.King] & friendlies
	kingMover := move.Mover(common.King)

	go func() {
		defer close(moveStream)

		for kings := allKings; kings != 0; kings &= kings - 1 {
			squareNum := common.FirstOne(kings)
			mover, capMover := kingMover(squareNum)

			moves := kingAttacks[squareNum]
			legalMoves := moves & ^friendlies

			for ; legalMoves != 0; legalMoves &= legalMoves - 1 {
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
