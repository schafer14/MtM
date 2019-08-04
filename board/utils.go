package board

func (b Board) opp() uint {
	if b.Turn == 0 {
		return 1
	}
	return 0
}

func (b Board) Opp() uint {
	if b.Turn == 0 {
		return 1
	}
	return 0
}

func (b Board) pieceOn(sq uint) (bool, uint) {
	if (b.Colors[0]|b.Colors[1])&(1<<sq) > 0 {
		for i := 0; i < 6; i++ {
			if b.Pieces[i]&(1<<sq) > 0 {
				return true, uint(i)
			}
		}
	}

	return false, 0
}
