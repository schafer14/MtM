package board

import (
	"fmt"
	"strings"

	"github.com/schafer14/chess/common"
)

// Create a board from a fen string
func fromFen(fen string) Board {
	if fen == "startpos" {
		return New()
	}

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
	col := n[0] - 97
	row := n[1] - 48

	return uint((row-1)*8 + col)
}

func (b Board) String() string {
	var fen string
	all := b.colors[common.White] | b.colors[common.Black]

	for row := 7; row >= 0; row-- {
		empty := 0
		for col := 0; col < 8; col++ {
			sqNum := 8*row + col
			square := uint64(1) << uint(sqNum)

			if all&square == 0 {
				empty += 1
			} else {
				if empty > 0 {
					fen += fmt.Sprintf("%v", empty)
					empty = 0
				}

				if square&b.colors[common.White]&b.pieces[common.Pawn] > 0 {
					fen += "P"
				}
				if square&b.colors[common.White]&b.pieces[common.Knight] > 0 {
					fen += "N"
				}
				if square&b.colors[common.White]&b.pieces[common.Bishop] > 0 {
					fen += "B"
				}
				if square&b.colors[common.White]&b.pieces[common.Rook] > 0 {
					fen += "R"
				}
				if square&b.colors[common.White]&b.pieces[common.Queen] > 0 {
					fen += "Q"
				}
				if square&b.colors[common.White]&b.pieces[common.King] > 0 {
					fen += "K"
				}
				if square&b.colors[common.Black]&b.pieces[common.Pawn] > 0 {
					fen += "p"
				}
				if square&b.colors[common.Black]&b.pieces[common.Knight] > 0 {
					fen += "n"
				}
				if square&b.colors[common.Black]&b.pieces[common.Bishop] > 0 {
					fen += "b"
				}
				if square&b.colors[common.Black]&b.pieces[common.Rook] > 0 {
					fen += "r"
				}
				if square&b.colors[common.Black]&b.pieces[common.Queen] > 0 {
					fen += "q"
				}
				if square&b.colors[common.Black]&b.pieces[common.King] > 0 {
					fen += "k"
				}
			}
		}
		if empty > 0 {
			fen += fmt.Sprintf("%v", empty)
		}
		if row > 0 {
			fen += "/"
		}
	}

	if b.turn == common.White {
		fen += " w "
	} else {
		fen += " b "
	}

	if b.castling[0] {
		fen += "K"
	}
	if b.castling[1] {
		fen += "Q"
	}
	if b.castling[2] {
		fen += "k"
	}
	if b.castling[3] {
		fen += "q"
	}

	if !b.castling[0] && !b.castling[1] && !b.castling[2] && !b.castling[3] {
		fen += "-"
	}

	fen += " "

	if b.enPassant == 0 {
		fen += "-"
	} else {
		square := b.enPassant
		row := square / 8
		col := square % 8

		fen += fmt.Sprintf("%s%v", cols[col], row+1)
	}

	fen += " 0 1"

	return fen
}
