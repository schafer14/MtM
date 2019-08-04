package board

import (
	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

// Move updates the state of the board given a move.
func (b *Board) Move(move move.Move32) (prev Board) {
	// Init variables
	opponent := b.opp()
	src := uint64(1 << move.Src())
	dest := uint64(1 << move.Dest())
	piece := move.Piece()
	isCap := move.IsCap()
	prev = *b

	// Clear old settings
	b.enPassant = 0

	// Update src and destination
	b.Pieces[piece] ^= src
	b.Colors[b.Turn] ^= src
	b.Pieces[piece] ^= dest
	b.Colors[b.Turn] ^= dest

	// Move rook if the move is a capture
	if isCastle, castleKing := move.Castle(); isCastle {
		if b.Turn == common.White && castleKing {
			b.Pieces[common.Rook] ^= 1 << 5
			b.Colors[common.White] ^= 1 << 5
			b.Pieces[common.Rook] ^= 1 << 7
			b.Colors[common.White] ^= 1 << 7
		}

		if b.Turn == common.White && !castleKing {
			b.Pieces[common.Rook] ^= 1 << 3
			b.Colors[common.White] ^= 1 << 3
			b.Pieces[common.Rook] ^= 1
			b.Colors[common.White] ^= 1
		}

		if b.Turn == common.Black && castleKing {
			b.Pieces[common.Rook] ^= 1 << 61
			b.Colors[common.Black] ^= 1 << 61
			b.Pieces[common.Rook] ^= 1 << 63
			b.Colors[common.Black] ^= 1 << 63
		}

		if b.Turn == common.Black && !castleKing {
			b.Pieces[common.Rook] ^= 1 << 59
			b.Colors[common.Black] ^= 1 << 59
			b.Pieces[common.Rook] ^= 1 << 56
			b.Colors[common.Black] ^= 1 << 56
		}
	}

	// Remove captured piece
	if isCap {
		capPiece, capSquare := move.Capture()
		b.Pieces[capPiece] ^= uint64(1 << capSquare)
		b.Colors[opponent] ^= uint64(1 << capSquare)
	}

	// Adds a promo piece
	if isPromo, promoPiece := move.Promotion(); isPromo {
		b.Pieces[piece] ^= dest
		b.Pieces[promoPiece] ^= dest
	}

	// Pawn en passant
	if piece == common.Pawn {
		// White en passant
		if move.Dest()-move.Src() == 16 {
			b.enPassant = move.Dest() - 8
		}

		// Black en passant
		if move.Src()-move.Dest() == 16 {
			b.enPassant = move.Dest() + 8
		}
	}

	// Remove castling privileges if king moves
	if piece == common.King {
		if b.Turn == common.White {
			b.castling[0] = false
			b.castling[1] = false
		}
		if b.Turn == common.Black {
			b.castling[2] = false
			b.castling[3] = false
		}
	}

	// Remove casting privleges on rook moves
	if move.Src() == 0 || move.Dest() == 0 {
		b.castling[1] = false
	}
	if move.Src() == 7 || move.Dest() == 7 {
		b.castling[0] = false
	}
	if move.Src() == 63 || move.Dest() == 63 {
		b.castling[2] = false
	}
	if move.Src() == 56 || move.Dest() == 56 {
		b.castling[3] = false
	}

	b.Turn = opponent

	return
}
