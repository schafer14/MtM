package main

import (
	"fmt"
	"time"

	"github.com/schafer14/chess/board"
	"github.com/schafer14/chess/move"
)

type moveCount struct {
	move  move.Move32
	count int
}

func main() {
	b := board.New()

	t := time.Now()
	x := Perft(b, 5)
	e := time.Since(t)
	fmt.Printf("perft: %v, time: %v\n", x, e)

	PrintDivide(b, 5)
}

func PrintDivide(b board.Board, depth int) {
	result := Divide(b, depth)
	count := 0
	for _, mc := range result {
		fmt.Printf("%s %v\n", mc.move.String(), mc.count)
		count += mc.count
	}

	fmt.Printf("\nMoves: %v\n", count)
}

func Divide(b board.Board, depth int) []moveCount {
	moves := make([]moveCount, 0, 256)

	for move := range b.MoveStream() {
		nb := b.Clone()
		nb.Move(move)
		moves = append(moves, moveCount{move: move, count: Perft(nb, depth-1)})
	}

	return moves
}

func Perft(b board.Board, depth int) int {
	if depth == 1 {
		return len(b.Moves())
	}

	count := 0
	for _, move := range b.Moves() {
		nb := b.Clone()
		nb.Move(move)
		count += Perft(nb, depth-1)
	}

	return count
}
