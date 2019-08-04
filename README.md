# MtM Chess Engine

MtM is a chess engine along with chess tools written in Golang. The engine is correct according
to perft testing of over one hundred positions. There is currently no agent/AI/uci interface as 
I am focusing on improving speed before working on uci.

Speed is currently being measured against the following position:

r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1

There is currently only 1 move generation function that is applied for all positions with a speed
of roughly 13 million moves per second on my laptop. You run this benchmark yourself using the
perft command.

## Future work

- Performance (with an aim of expanding 40-50 million nps on my laptop: comparable to Roce38)
- API UI: The API needs to be cleaned up quite a bit
- Documentation
- UCI/gameplaying agent


## Stability Warning

The API will change considerably for this project.
