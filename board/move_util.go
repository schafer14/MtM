package board

import (
	"errors"

	"github.com/schafer14/chess/common"
	"github.com/schafer14/chess/move"
)

var rowMap = map[byte]uint{
	'1': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
}

var colMap = map[byte]uint{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
}

var pieceMap = map[byte]uint{
	'N': common.Knight,
	'B': common.Bishop,
	'R': common.Rook,
	'Q': common.Queen,
}

// Returns a move object from a src dest move e2e4. Not working
// for chess 960
func (b Board) MoveFromSrcDestNotation(moveStr string) (move.Move32, error) {
	if len(moveStr) != 4 && len(moveStr) != 5 {
		return 0, errors.New("Invalid format: expected format e2e4")
	}

	srcCol, isOk := colMap[moveStr[0]]
	if !isOk {
		return 0, errors.New("Invalid format: expected format e2e4")
	}
	srcRow, isOk := rowMap[moveStr[1]]
	if !isOk {
		return 0, errors.New("Invalid format: expected format e2e4")
	}
	destCol, isOk := colMap[moveStr[2]]
	if !isOk {
		return 0, errors.New("Invalid format: expected format e2e4")
	}
	destRow, isOk := rowMap[moveStr[3]]
	if !isOk {
		return 0, errors.New("Invalid format: expected format e2e4")
	}

	srcSquare := srcRow*8 + srcCol
	destSquare := destRow*8 + destCol
	hasPiece, piece := b.pieceOn(srcSquare)
	if !hasPiece {
		return 0, errors.New("Illegal move")
	}
	isCap, capPiece := b.pieceOn(destSquare)

	move := move.New(piece, srcSquare, destSquare)

	if isCap {
		move = move.SetCap(destSquare, capPiece)
	}

	if piece == common.Pawn && b.turn == common.White && b.enPassant == destSquare-8 {
		isCap, capPiece = b.pieceOn(destSquare - 8)
		move = move.SetCap(destSquare-8, capPiece)
	}

	if piece == common.Pawn && b.turn == common.Black && b.enPassant == destSquare+8 {
		isCap, capPiece = b.pieceOn(destSquare + 8)
		move = move.SetCap(destSquare+8, capPiece)
	}

	if len(moveStr) == 5 {
		promoPiece := pieceMap[moveStr[4]]
		move = move.SetPromo(promoPiece)
	}

	if piece == common.King && srcSquare == 4 && destSquare == 6 {
		move = move.CastleKing()
	}
	if piece == common.King && srcSquare == 4 && destSquare == 2 {
		move = move.CastleQueen()
	}
	if piece == common.King && srcSquare == 60 && destSquare == 62 {
		move = move.CastleKing()
	}
	if piece == common.King && srcSquare == 60 && destSquare == 58 {
		move = move.CastleQueen()
	}

	moves := b.Moves()
	isLegal := false
	for {
		hasNext, legalMove := moves.Next()
		if !hasNext {
			break
		}
		if uint32(move) == uint32(legalMove) {
			isLegal = true
			break
		}
	}

	if !isLegal {
		return 0, errors.New("Illegal move")
	}

	return move, nil
}
