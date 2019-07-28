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
	prev = b.Clone()

	// Clear old settings
	b.enPassant = 0

	// Update src and destination
	b.pieces[piece] ^= src
	b.colors[b.turn] ^= src
	b.pieces[piece] ^= dest
	b.colors[b.turn] ^= dest

	// Move rook if the move is a capture
	if isCastle, castleKing := move.Castle(); isCastle {
		if b.turn == common.White && castleKing {
			b.pieces[common.Rook] ^= 1 << 5
			b.colors[common.White] ^= 1 << 5
			b.pieces[common.Rook] ^= 1 << 7
			b.colors[common.White] ^= 1 << 7
		}

		if b.turn == common.White && !castleKing {
			b.pieces[common.Rook] ^= 1 << 3
			b.colors[common.White] ^= 1 << 3
			b.pieces[common.Rook] ^= 1
			b.colors[common.White] ^= 1
		}

		if b.turn == common.Black && castleKing {
			b.pieces[common.Rook] ^= 1 << 61
			b.colors[common.Black] ^= 1 << 61
			b.pieces[common.Rook] ^= 1 << 63
			b.colors[common.Black] ^= 1 << 63
		}

		if b.turn == common.Black && !castleKing {
			b.pieces[common.Rook] ^= 1 << 59
			b.colors[common.Black] ^= 1 << 59
			b.pieces[common.Rook] ^= 1 << 56
			b.colors[common.Black] ^= 1 << 56
		}
	}

	// Remove captured piece
	if isCap {
		capPiece, capSquare := move.Capture()
		b.pieces[capPiece] ^= uint64(1 << capSquare)
		b.colors[opponent] ^= uint64(1 << capSquare)
	}

	// Adds a promo piece
	if isPromo, promoPiece := move.Promotion(); isPromo {
		b.pieces[piece] ^= dest
		b.pieces[promoPiece] ^= dest
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
		if b.turn == common.White {
			b.castling[0] = false
			b.castling[1] = false
		}
		if b.turn == common.Black {
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

	b.turn = opponent

	return
}
