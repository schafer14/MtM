package move

import (
	"fmt"

	"github.com/schafer14/chess/common"
)

// From least to most significant bytes
// 0-6 : src square
// 7-13 : dest square
// 14-16 : piece
// 17-19 : captured piece
// 20-25 : captured square
// 26 : is capture
// 27-29 : promotion piece
// 30 : king side castle
// 31 : queen side castle
type Move32 uint32

// Mover incrementally produces a move function.
// `
// pawnMoves := Generate(common.Pawn)
// for pawn := range pawns {
//		mover := pawnMoves(pawn.src)
//    for moves := range pawn.moves {
//      move := moves(move.dest, common.Knight)
//    }
// }
func Mover(piece uint) func(uint) (func(uint) Move32, func(uint, uint) Move32) {
	return func(src uint) (func(uint) Move32, func(uint, uint) Move32) {
		return func(dest uint) Move32 {
				return Move32(src | dest<<6 | piece<<13)
			}, func(dest uint, capPiece uint) Move32 {
				return Move32(src | dest<<6 | piece<<13 | capPiece<<16 | dest<<19 | 1<<25)
			}
	}
}

// New creates the basis for a new move
func New(piece uint, src uint, dest uint) Move32 {
	return Move32(src | dest<<6 | piece<<13)
}

func PawnQuiet(src uint, dest uint) Move32 {
	return Move32(src | dest<<6 | common.Pawn<<13)
}

func (m Move32) SetCap(capSquare uint, capPiece uint) Move32 {
	return Move32(uint(m) | capPiece<<16 | capSquare<<19 | 1<<25)
}

func (m Move32) SetCastleKing() Move32 {
	return Move32(uint(m) | 1<<29)
}

func (m Move32) SetCastleQueen() Move32 {
	return Move32(uint(m) | 1<<30)
}

func (m Move32) SetPromo(piece uint) Move32 {
	return Move32(uint(m) | piece<<26)
}

func (m Move32) CastleKing() Move32 {
	return Move32(uint(m) | 1<<29)
}

func (m Move32) CastleQueen() Move32 {
	return Move32(uint(m) | 1<<30)
}

func (m Move32) Src() uint {
	return uint(m) & 0x3F
}

func (m Move32) Dest() uint {
	return uint(m) >> 6 & 0x3F
}

func (m Move32) Piece() uint {
	return uint(m) >> 13 & 0x07
}

func (m Move32) Capture() (uint, uint) {
	return uint(m) >> 16 & 0x07, uint(m) >> 19 & 0x3F
}

func (m Move32) IsCap() bool {
	return uint(m)>>25&0x01 > 0
}

func (m Move32) Promotion() (bool, uint) {
	return uint(m)>>26&0x07 > 0, uint(m) >> 26 & 0x07
}

func (m Move32) Castle() (bool, bool) {
	return uint(m)>>29&0x03 > 0, uint(m)>>29&0x01 > 0
}

var roce38StyleMoves bool = false

func SetRoceMoves(isRoce bool) {
	roce38StyleMoves = isRoce
}

func (m Move32) String() string {
	chars := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	pieces := []rune{'P', 'N', 'B', 'R', 'Q', 'K'}

	if isPromo, promoPiece := m.Promotion(); isPromo && !roce38StyleMoves {
		return fmt.Sprintf(
			"%c%v%c%v%c",
			chars[m.Src()%8],
			m.Src()/8+1,
			chars[m.Dest()%8],
			m.Dest()/8+1,
			pieces[promoPiece],
		)
	} else {
		return fmt.Sprintf(
			"%c%v%c%v",
			chars[m.Src()%8],
			m.Src()/8+1,
			chars[m.Dest()%8],
			m.Dest()/8+1,
		)
	}
}
