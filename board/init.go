package board

import (
	"math/bits"
)

var rookMagicMoves [64][4096]uint64
var bishopMagicMoves [64][4096]uint64

type Magic struct {
	mask  uint64
	magic uint64
}

func init() {
	initMagic()
}

func initMagic() {
	for i := 0; i < 64; i++ {
		buildRookMagic(uint(i))
		buildBishopMagic(uint(i))
	}
}

func buildRookMagic(square uint) {
	magic := rookMagic[square].magic
	mask := rookMagic[square].mask
	bits := uint(52)
	var squareBB uint64 = 1 << square

	permutations := getPermutations(0, mask)

	for _, blocker := range permutations {
		index := (blocker * magic) >> bits
		rookMagicMoves[square][index] = straightBB(blocker, squareBB)
	}
}

func buildBishopMagic(square uint) {
	magic := bishopMagic[square].magic
	mask := bishopMagic[square].mask
	bits := uint(55)
	var squareBB uint64 = 1 << square

	permutations := getPermutations(0, mask)

	for _, blocker := range permutations {
		index := (blocker * magic) >> bits
		bishopMagicMoves[square][index] = diagBB(blocker, squareBB)
	}
}

func getPermutations(set uint64, mutable uint64) []uint64 {
	if bits.OnesCount64(mutable) == 0 {
		return []uint64{set}
	}

	bit := mutable & -mutable
	mutable ^= bit

	withBitSet := getPermutations(set|bit, mutable)
	withoutBitSet := getPermutations(set, mutable)

	return append(withBitSet, withoutBitSet...)
}

func straightBB(occ uint64, square uint64) uint64 {
	squareNum := bits.TrailingZeros64(square)

	forward := slideAttacks(occ, square, columns[squareNum%8])
	right := slideAttacks(occ, square, ranks[squareNum/8])
	backwards := reversSlideAttacks(occ, square, columns[squareNum%8])
	left := reversSlideAttacks(occ, square, ranks[squareNum/8])

	return forward | right | backwards | left
}

/*
	Generates a bitboard containing all the legal straight moves.
*/
func diagBB(occ uint64, square uint64) uint64 {
	squareNum := bits.TrailingZeros64(square)

	mask := diag[((squareNum/8)-(squareNum%8))&15]
	antiMask := antiDiag[7^((squareNum/8)+(squareNum%8))]

	northEast := slideAttacks(occ, square, mask)
	northWest := slideAttacks(occ, square, antiMask)
	southWest := reversSlideAttacks(occ, square, mask)
	southEast := reversSlideAttacks(occ, square, antiMask)

	return northEast | southWest | northWest | southEast
}

/*
	Generates move bitboard for sliding pieces using positive rays
*/
func slideAttacks(occ uint64, square uint64, mask uint64) uint64 {
	potentialBlockers := occ & mask

	diff := potentialBlockers - 2*square
	changed := diff ^ occ

	return changed & mask
}

/*
	Generates move bitboard for sliding pieces using negitive rays
*/
func reversSlideAttacks(occ uint64, square uint64, maskB uint64) uint64 {
	o := bits.Reverse64(occ)
	s := bits.Reverse64(square)
	mask := bits.Reverse64(maskB)

	potentialBlockers := o & mask

	diff := potentialBlockers - 2*s
	changed := diff ^ o

	return bits.Reverse64(changed & mask)
}
