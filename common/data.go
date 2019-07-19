package common

const (
	White uint = iota
	Black
)

const (
	Pawn uint = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

const (
	ColA uint64 = 0x0101010101010101 << iota
	ColB
	ColC
	ColD
	ColE
	ColF
	ColG
	ColH
)

const (
	Row1 uint64 = 0xFF << (8 * iota)
	Row2
	Row3
	Row4
	Row5
	Row6
	Row7
	Row8
)
