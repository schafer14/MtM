package board

func (b Board) opp() uint {
	if b.turn == 0 {
		return 1
	}
	return 0
}

func (b Board) pieceOn(sq uint) (bool, uint) {
	if (b.colors[0]|b.colors[1])&(1<<sq) > 0 {
		for i := 0; i < 6; i++ {
			if b.pieces[i]&(1<<sq) > 0 {
				return true, uint(i)
			}
		}
	}

	return false, 0
}
