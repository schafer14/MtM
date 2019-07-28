package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

func promote(move move.Move32, moves *[]move.Move32) {
	*moves = append(*moves, move.SetPromo(common.Queen))
	*moves = append(*moves, move.SetPromo(common.Knight))
	*moves = append(*moves, move.SetPromo(common.Rook))
	*moves = append(*moves, move.SetPromo(common.Bishop))
}

func (b Board) pawnMoves(moves *[]move.Move32) {
	opp := b.colors[b.opp()] | (1 << b.enPassant)
	friendlies := b.colors[b.turn]
	pawns := b.pieces[common.Pawn] & friendlies
	all := b.colors[0] | b.colors[1]
	empty := ^all

	// White
	if b.turn == common.White {
		// Forward 1
		for sourceBB := pawns & (empty >> 8); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source+8 >= 56 {
				promote(move.PawnQuiet(source, source+8), moves)
			} else {
				*moves = append(*moves, move.PawnQuiet(source, source+8))
			}
		}

		// Forward 2
		for sourceBB := pawns & (empty >> 8) & (empty >> 16) & common.Row2; sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			*moves = append(*moves, move.PawnQuiet(source, source+16))
		}

		// Cap Right
		for sourceBB := pawns & (^common.ColH) & (opp >> 9); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source+9 == b.enPassant {
				*moves = append(*moves, move.PawnQuiet(source, source+9).SetCap(source+1, common.Pawn))
			} else {
				_, piece := b.pieceOn(source + 9)
				if source+9 >= 56 {
					promote(move.PawnQuiet(source, source+9).SetCap(source+9, piece), moves)
				} else {
					*moves = append(*moves, move.PawnQuiet(source, source+9).SetCap(source+9, piece))
				}
			}
		}

		// Cap Left
		for sourceBB := pawns & (^common.ColA) & (opp >> 7); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source+7 == b.enPassant {
				*moves = append(*moves, move.PawnQuiet(source, source+7).SetCap(source-1, common.Pawn))
			} else {
				_, piece := b.pieceOn(source + 7)
				if source+7 >= 56 {
					promote(move.PawnQuiet(source, source+7).SetCap(source+7, piece), moves)
				} else {
					*moves = append(*moves, move.PawnQuiet(source, source+7).SetCap(source+7, piece))
				}
			}
		}
	}

	// Black
	if b.turn == common.Black {
		// Forward 1
		for sourceBB := pawns & (empty << 8); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source-8 <= 7 {
				promote(move.PawnQuiet(source, source-8), moves)
			} else {
				*moves = append(*moves, move.PawnQuiet(source, source-8))
			}
		}

		// Forward 2
		for sourceBB := pawns & (empty << 8) & (empty << 16) & common.Row7; sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			*moves = append(*moves, move.PawnQuiet(source, source-16))
		}

		// Cap Left
		for sourceBB := pawns & (^common.ColH) & (opp << 7); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source-7 == b.enPassant {
				_, piece := b.pieceOn(source + 1)
				*moves = append(*moves, move.PawnQuiet(source, source-7).SetCap(source+1, piece))
			} else {
				_, piece := b.pieceOn(source - 7)
				if source-7 <= 7 {
					promote(move.PawnQuiet(source, source-7).SetCap(source-7, piece), moves)
				} else {
					*moves = append(*moves, move.PawnQuiet(source, source-7).SetCap(source-7, piece))
				}
			}
		}

		// Cap Right
		for sourceBB := pawns & (^common.ColA) & (opp << 9); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source-9 == b.enPassant {
				_, piece := b.pieceOn(source - 1)
				*moves = append(*moves, move.PawnQuiet(source, source-9).SetCap(source-1, piece))
			} else {
				_, piece := b.pieceOn(source - 9)
				if source-9 <= 7 {
					promote(move.PawnQuiet(source, source-9).SetCap(source-9, piece), moves)
				} else {
					*moves = append(*moves, move.PawnQuiet(source, source-9).SetCap(source-9, piece))
				}
			}
		}
	}

}

func (b Board) pawnAttacks(turn uint) (attackSpace uint64) {
	pawns := b.pieces[common.Pawn] & b.colors[turn]

	if turn == common.Black {
		attackSpace |= (pawns >> 9 & ^common.ColH)
		attackSpace |= (pawns >> 7 & ^common.ColA)
	}

	if turn == common.White {
		attackSpace |= (pawns << 7 & ^common.ColH)
		attackSpace |= (pawns << 9 & ^common.ColA)
	}

	return attackSpace
}
