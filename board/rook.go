package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func (b Board) rookMoves() <-chan move.Move32 {
	moveStream := make(chan move.Move32)
	occ := b.colors[0] | b.colors[1]
	friendlies := b.colors[b.turn]
	allRooks := b.pieces[common.Rook] & friendlies
	rookMover := move.Mover(common.Rook)

	go func() {
		defer close(moveStream)

		for rooks := allRooks; rooks != 0; rooks &= rooks - 1 {
			squareNum := common.FirstOne(rooks)
			mover, capMover := rookMover(squareNum)

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

		}
	}()

	return moveStream
}
