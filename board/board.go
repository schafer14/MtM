package board

import (
	"github.com/schafer14/chess/move"
)

// Moves generates a list of legal moves in a position
func (b Board) Moves() []move.Move32 {
	moves := make([]move.Move32, 0, 256)

	b.pawnMoves(&moves)
	b.knightMoves(&moves)
	b.bishopMoves(&moves)
	b.rookMoves(&moves)
	b.queenMoves(&moves)
	b.kingMoves(&moves)

	return moves
}

// Undo reverts the board to the state previous to the move.
func (b *Board) Undo(move move.Move32) {}

// Clone duplicates a chess board.
func (b Board) Clone() Board {
	newBoard := b
	return newBoard
}

// New creates a new board initialized to the intial position.
func New() Board {
	return initialBoard.Clone()
}

// Empty creates a board with no pieces on it.
func Empty() Board {
	return Board{}
}

// FromFen creates a new board from a fenstring.
var FromFen = fromFen

// ReadMove reads a string input into a move given a board position
func (b Board) ReadMove(moveString string) move.Move32 {
	return move.Move32(0)
}
