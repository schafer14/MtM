package main

import (
	"testing"
	"time"

	"github.com/schafer14/MtM/board"
)

func TestPeft(t *testing.T) {
	t1 := time.Now()
	x := Perft(board.FromFen("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"), 5)
	e := time.Since(t1)
	t.Logf("perft: %v, time: %v\nn/s: %f\n", x, e, float64(x)/e.Seconds())

}
