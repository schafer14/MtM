.PHONY: perft
perft:
	go build -o perft-test ./perft && ./perft-test && rm ./perft-test
