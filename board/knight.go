package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func (b Board) knightMoves() <-chan move.Move32 {
	moveStream := make(chan move.Move32)
	friendlies := b.colors[b.turn]
	allKnights := b.pieces[common.Knight] & friendlies
	knightMover := move.Mover(common.Knight)

	go func() {
		defer close(moveStream)

		for knights := allKnights; knights != 0; knights &= knights - 1 {
			squareNum := common.FirstOne(knights)
			mover, capMover := knightMover(squareNum)

			moves := knightAttacks[squareNum]
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
