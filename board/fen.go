package board

import (
	"strings"

	"github.com/schafer14/chess/common"
)

// Create a board from a fen string
func fromFen(fen string) Board {
	board := Board{}
	parts := strings.Split(fen, " ")
	pos := parts[0]
	rows := strings.Split(pos, "/")

	// Add pieces and white black to board
	for i, row := range rows {
		n := 0
		for _, char := range row {
			squareNum := uint((7-i)*8 + n)
			square := uint64(1) << squareNum
			if char >= '0' && char <= '9' {
				n += int(char) - 48
				continue
			}
			if char >= 'A' && char <= 'Z' {
				board.colors[common.White] |= square
			}
			if char >= 'a' && char <= 'z' {
				board.colors[common.Black] |= square
			}
			if char == 'p' || char == 'P' {
				board.pieces[common.Pawn] |= square
			}
			if char == 'n' || char == 'N' {
				board.pieces[common.Knight] |= square
			}
			if char == 'b' || char == 'B' {
				board.pieces[common.Bishop] |= square
			}
			if char == 'r' || char == 'R' {
				board.pieces[common.Rook] |= square
			}
			if char == 'q' || char == 'Q' {
				board.pieces[common.Queen] |= square
			}
			if char == 'k' || char == 'K' {
				board.pieces[common.King] |= square
			}
			n++
		}
	}

	// Set color
	if parts[1] == "b" {
		board.turn = common.Black
	}

	// castle privileges
	if strings.ContainsRune(parts[2], 'K') {
		board.castling[0] = true
	}
	if strings.ContainsRune(parts[2], 'Q') {
		board.castling[1] = true
	}
	if strings.ContainsRune(parts[2], 'k') {
		board.castling[2] = true
	}
	if strings.ContainsRune(parts[2], 'q') {
		board.castling[3] = true
	}

	// enPassant
	if parts[3] != "-" {
		board.enPassant = squareToNum(parts[3])
	}

	return board
}

func squareToNum(n string) uint {
	col := n[0] - 96
	row := n[1] - 48

	return uint((row-1)*8 + col)
}
