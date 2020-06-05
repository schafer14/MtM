package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/schafer14/MtM/board"
	"github.com/schafer14/MtM/move"
	"github.com/urfave/cli"
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
	app := cli.NewApp()
	app.Name = "MtM Perft Tool"
	app.Usage = "perft, test"
	app.Commands = []cli.Command{
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Run the test suite",
			Action: func(c *cli.Context) error {
				P()
				return nil
			},
		},
		{
			Name:    "moves",
			Aliases: []string{"m"},
			Usage:   "Returns a list of moves in a position",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "pos, p",
					Value: "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
					Usage: "Position",
				},
			},
			Action: func(c *cli.Context) error {
				b := board.FromFen(c.String("pos"))
				ml := b.Moves()
				for {
					hasNext, move := ml.Next()

					if !hasNext {
						break
					}

					fmt.Printf("%+v ", move)

				}
				return nil
			},
		},
		{
			Name:    "pos",
			Aliases: []string{"p"},
			Usage:   "Run the test on a single position",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "pos, p",
					Value: "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
					Usage: "Position to run tests against",
				},
				cli.IntFlag{
					Name:  "depth, d",
					Value: 4,
					Usage: "Depth to test to",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Printf("board = %+v\ndepth = %+v\n", c.String("pos"), c.Int("depth"))
				t1 := time.Now()
				x := Perft(board.FromFen(c.String("pos")), c.Int("depth"))
				e := time.Since(t1)
				fmt.Printf("expanded = %+v\ntime: %+v\nn/s: %+v\n", x, e, float64(x)/e.Seconds())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

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
	var ml board.MoveList
	var nb board.Board
	b.PsudoMoves(&ml)

	for {
		hasNext, move := ml.Next()
		if !hasNext {
			break
		}
		nb = b
		nb.Move(move)
		if nb.IsInCheck(nb.Opp()) {
			continue
		}
		moves = append(moves, moveCount{move: move, count: Perft(nb, depth-1)})
	}

	return moves
}

func Perft(b board.Board, depth int) int {
	if depth < 1 {
		return 1
	}

	if depth == 1 {
		return b.Moves().Len()
	}

	count := 0
	var ml board.MoveList
	var nb board.Board
	b.PsudoMoves(&ml)
	for {
		hasNext, move := ml.Next()
		if !hasNext {
			break
		}
		nb = b
		nb.Move(move)
		if nb.IsInCheck(nb.Opp()) {
			continue
		}
		count += Perft(nb, depth-1)
	}

	return count
}
