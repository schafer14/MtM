package util

import (
	"sync"

	"github.com/schafer14/chess/move"
)

func MergeMoveStreams(streams ...<-chan move.Move32) <-chan move.Move32 {
	var wg sync.WaitGroup
	multiplexedStream := make(chan move.Move32)

	multiplex := func(ch <-chan move.Move32) {
		defer wg.Done()

		for m := range ch {
			multiplexedStream <- m
		}
	}

	wg.Add(len(streams))
	for _, c := range streams {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
