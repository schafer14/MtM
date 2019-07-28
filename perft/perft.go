package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/profile"
	"github.com/schafer14/chess/board"
	"github.com/schafer14/chess/move"
)

const STARTPOS = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type moveCount struct {
	move  move.Move32
	count int
}

type moveStringCount struct {
	move  string
	count int
}

type divideOutput struct {
	moveStr []moveStringCount
	moves   int
	nodes   int
}

func main() {
	// P()

	defer profile.Start(profile.CPUProfile).Stop()
	t := time.Now()
	x := Perft(board.FromFen("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"), 5)
	e := time.Since(t)
	fmt.Printf("perft: %v, time: %v\nn/s: %f\n", x, e, float64(x)/e.Seconds())
	// errors := roceCmp(position{fen: "startpos"}, 6)

	// for _, err := range errors {
	// 	fmt.Println(err)
	// }

	// if len(errors) == 0 {
	// 	fmt.Println("No errors")
	// }
}

func (de divideOutput) String() (str string) {

	for _, move := range de.moveStr {
		str += fmt.Sprintf("%v %v\n", move.move, move.count)
	}

	str += fmt.Sprintf("\nMoves %v\n", de.moves)
	str += fmt.Sprintf("Nodes %v\n", de.nodes)
	return str
}

func P() {
	f, err := os.Open("./perft.epd")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	buf := bufio.NewReader(f)

	for {
		x, err := buf.ReadString('\n')

		if err != nil {
			return
		}

		move.SetRoceMoves(true)
		parts := strings.Split(x, ";")
		fmt.Printf("Testing %v\n", parts[0])

		errors := roceCmp(position{fen: parts[0]}, len(parts)-2)

		for _, err := range errors {
			fmt.Println(err)
			return
		}

		errors = roceCmp(position{fen: parts[0]}, len(parts)-1)

		for _, err := range errors {
			fmt.Println(err)
			return
		}

		if len(errors) == 0 {
			fmt.Printf("No errors for position %v\n", parts[0])
		}
	}
}

func PerftDivide(fen string, d int) {
	var b board.Board

	if fen == "startpos" {
		b = board.New()
	} else {
		b = board.FromFen(fen)
	}

	t := time.Now()
	x := Perft(b, d)
	e := time.Since(t)
	fmt.Printf("perft: %v, time: %v\n", x, e)
}

func divide(b board.Board, depth int) divideOutput {
	result := divideRecursive(b, depth)
	count := 0
	moves := []moveStringCount{}
	for _, mc := range result {
		moves = append(moves, moveStringCount{count: mc.count, move: mc.move.String()})
		count += mc.count
	}

	return divideOutput{
		moveStr: moves,
		moves:   len(result),
		nodes:   count,
	}
}

func divideRecursive(b board.Board, depth int) []moveCount {
	moves := make([]moveCount, 0, 256)

	for _, move := range b.Moves() {
		nb := b.Clone()
		nb.Move(move)
		moves = append(moves, moveCount{move: move, count: Perft(nb, depth-1)})
	}

	return moves
}

func Perft(b board.Board, depth int) int {
	if depth < 1 {
		return 1
	}

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
