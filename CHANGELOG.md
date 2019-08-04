* Changed piece moves signiture from `func () chan <- Move32` to `func(moves *[]Move32)` which improved the speed of perft 5 from ~3s to ~150ms.

* Changed *[]Move32 to a MoveList to stop allocation

* Combination of perft testing and pprof to remove some inefficencies
  * Moved from 6 million nodes per second to 13 million in perft testing
