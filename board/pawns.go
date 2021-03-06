package board

import (
	"github.com/schafer14/MtM/common"
	"github.com/schafer14/MtM/move"
)

func promote(move move.Move32, moves *MoveList) {
	moves.Append(move.SetPromo(common.Queen))
	moves.Append(move.SetPromo(common.Knight))
	moves.Append(move.SetPromo(common.Rook))
	moves.Append(move.SetPromo(common.Bishop))
}

func (b Board) pawnMoves(moves *MoveList) {
	opp := b.Colors[b.opp()] | (1 << b.enPassant)
	friendlies := b.Colors[b.Turn]
	pawns := b.Pieces[common.Pawn] & friendlies
	all := b.Colors[0] | b.Colors[1]
	empty := ^all

	// White
	if b.Turn == common.White {
		// Forward 1
		for sourceBB := pawns & (empty >> 8); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source+8 >= 56 {
				promote(move.PawnQuiet(source, source+8), moves)
			} else {
				moves.Append(move.PawnQuiet(source, source+8))
			}
		}

		// Forward 2
		for sourceBB := pawns & (empty >> 8) & (empty >> 16) & common.Row2; sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			moves.Append(move.PawnQuiet(source, source+16))
		}

		// Cap Right
		for sourceBB := pawns & (^common.ColH) & (opp >> 9); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source+9 == b.enPassant {
				moves.Append(move.PawnQuiet(source, source+9).SetCap(source+1, common.Pawn))
			} else {
				_, piece := b.pieceOn(source + 9)
				if source+9 >= 56 {
					promote(move.PawnQuiet(source, source+9).SetCap(source+9, piece), moves)
				} else {
					moves.Append(move.PawnQuiet(source, source+9).SetCap(source+9, piece))
				}
			}
		}

		// Cap Left
		for sourceBB := pawns & (^common.ColA) & (opp >> 7); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source+7 == b.enPassant {
				moves.Append(move.PawnQuiet(source, source+7).SetCap(source-1, common.Pawn))
			} else {
				_, piece := b.pieceOn(source + 7)
				if source+7 >= 56 {
					promote(move.PawnQuiet(source, source+7).SetCap(source+7, piece), moves)
				} else {
					moves.Append(move.PawnQuiet(source, source+7).SetCap(source+7, piece))
				}
			}
		}
	}

	// Black
	if b.Turn == common.Black {
		// Forward 1
		for sourceBB := pawns & (empty << 8); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source-8 <= 7 {
				promote(move.PawnQuiet(source, source-8), moves)
			} else {
				moves.Append(move.PawnQuiet(source, source-8))
			}
		}

		// Forward 2
		for sourceBB := pawns & (empty << 8) & (empty << 16) & common.Row7; sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			moves.Append(move.PawnQuiet(source, source-16))
		}

		// Cap Left
		for sourceBB := pawns & (^common.ColH) & (opp << 7); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source-7 == b.enPassant {
				_, piece := b.pieceOn(source + 1)
				moves.Append(move.PawnQuiet(source, source-7).SetCap(source+1, piece))
			} else {
				_, piece := b.pieceOn(source - 7)
				if source-7 <= 7 {
					promote(move.PawnQuiet(source, source-7).SetCap(source-7, piece), moves)
				} else {
					moves.Append(move.PawnQuiet(source, source-7).SetCap(source-7, piece))
				}
			}
		}

		// Cap Right
		for sourceBB := pawns & (^common.ColA) & (opp << 9); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source-9 == b.enPassant {
				_, piece := b.pieceOn(source - 1)
				moves.Append(move.PawnQuiet(source, source-9).SetCap(source-1, piece))
			} else {
				_, piece := b.pieceOn(source - 9)
				if source-9 <= 7 {
					promote(move.PawnQuiet(source, source-9).SetCap(source-9, piece), moves)
				} else {
					moves.Append(move.PawnQuiet(source, source-9).SetCap(source-9, piece))
				}
			}
		}
	}

}

func (b Board) pawnAttacks(turn uint) (attackSpace uint64) {
	pawns := b.Pieces[common.Pawn] & b.Colors[turn]

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
