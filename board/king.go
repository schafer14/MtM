package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func (b Board) kingMoves(movesSlice *[]move.Move32) {
	friendlies := b.colors[b.turn]
	allKings := b.pieces[common.King] & friendlies
	kingMover := move.Mover(common.King)
	oppAttack := b.attackSpace(b.opp())
	all := (friendlies | b.colors[b.opp()])

	for kings := allKings; kings != 0; kings &= kings - 1 {
		squareNum := common.FirstOne(kings)
		mover, capMover := kingMover(squareNum)

		moves := kingAttacks[squareNum]
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

	// Castle White does not work for chess 960
	if b.turn == common.White && b.colors[common.White]&b.pieces[common.King] == 0x10 {
		if 0x70&oppAttack == 0 && b.castling[0] && all&0x60 == 0 {
			*movesSlice = append(*movesSlice, move.New(common.King, 4, 6).SetCastleKing())
		}

		if 0x1c&oppAttack == 0 && b.castling[1] && all&0x0e == 0 {
			*movesSlice = append(*movesSlice, move.New(common.King, 4, 2).SetCastleQueen())
		}
	}

	if b.turn == common.Black && b.colors[common.Black]&b.pieces[common.King] == 0x1000000000000000 {
		if 0x7000000000000000&oppAttack == 0 && b.castling[2] && 0x6000000000000000&all == 0 {
			*movesSlice = append(*movesSlice, move.New(common.King, 60, 62).SetCastleKing())
		}

		if 0x1c00000000000000&oppAttack == 0 && b.castling[3] && 0x0e00000000000000&all == 0 {
			*movesSlice = append(*movesSlice, move.New(common.King, 60, 58).SetCastleQueen())
		}
	}

}

func (b Board) kingAttacks(turn uint) (attackSpace uint64) {
	friendlies := b.colors[turn]
	allKings := b.pieces[common.King] & friendlies

	for kings := allKings; kings != 0; kings &= kings - 1 {
		squareNum := common.FirstOne(kings)

		attackSpace |= kingAttacks[squareNum]
	}

	return attackSpace
}
