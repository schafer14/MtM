package board

import (
	"github.com/schafer14/chess/move"
	"github.com/schafer14/chess/util"
)

// MoveSteam returns a channel of moves that contains all the
// legal moves in a given position.
func (b Board) MoveStream() <-chan move.Move32 {
	return util.MergeMoveStreams(
		b.pawnMoves(),
		b.knightMoves(),
		b.bishopMoves(),
		b.rookMoves(),
		b.queenMoves(),
		b.kingMoves(),
	)
}

func (b Board) Moves() (moves []move.Move32) {
	moveStream := b.MoveStream()
	for move := range moveStream {
		moves = append(moves, move)
	}
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
