package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

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
			*moves = append(*moves, move.PawnQuiet(source, source+8))
		}

		// Forward 2
		for sourceBB := pawns & (empty >> 8) & (empty >> 16) & common.Row2; sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			*moves = append(*moves, move.PawnQuiet(source, source+16))
		}

		// Cap Left
		for sourceBB := pawns & (^common.ColH) & (opp >> 9); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source+9 == b.enPassant {
				_, piece := b.pieceOn(source + 1)
				*moves = append(*moves, move.PawnQuiet(source, source+9).SetCap(source+1, piece))
			} else {
				_, piece := b.pieceOn(source + 9)
				*moves = append(*moves, move.PawnQuiet(source, source+9).SetCap(source+9, piece))
			}
		}

		// Cap Right
		for sourceBB := pawns & (^common.ColA) & (opp >> 7); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			if source+7 == b.enPassant {
				_, piece := b.pieceOn(source - 1)
				*moves = append(*moves, move.PawnQuiet(source, source+9).SetCap(source-1, piece))
			} else {
				_, piece := b.pieceOn(source + 7)
				*moves = append(*moves, move.PawnQuiet(source, source+7).SetCap(source+7, piece))
			}
		}
	}

	// Black
	if b.turn == common.Black {
		// Forward 1
		for sourceBB := pawns & (empty << 8); sourceBB != 0; sourceBB &= sourceBB - 1 {
			source := common.FirstOne(sourceBB)
			*moves = append(*moves, move.PawnQuiet(source, source-8))
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
				*moves = append(*moves, move.PawnQuiet(source, source-7).SetCap(source-7, piece))
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
				*moves = append(*moves, move.PawnQuiet(source, source-9).SetCap(source-9, piece))
			}
		}
	}

}
