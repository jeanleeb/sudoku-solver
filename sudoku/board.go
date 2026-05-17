package sudoku

type Board [9][9]int

func NewBoard() *Board {
	b := Board{}
	return &b
}

func (b *Board) Get(row, col int) int {
	return b[row][col]
}

func (b *Board) Set(row, col, value int) {
	b[row][col] = value
}

func getBox(row, col int) (int, int) {
	startRow, startCol := (row/3)*3, (col/3)*3

	return startRow, startCol
}

func (b *Board) IsValidPlacement(row, col, val int) bool {
	if val == 0 {
		return true
	}

	for j, v := range b[row] {
		if j == col {
			continue
		}
		if v == val {
			return false
		}
	}

	for i, r := range b {
		if i == row {
			continue
		}
		if r[col] == val {
			return false
		}
	}

	boxStartRow, boxStartCol := getBox(row, col)
	for i := boxStartRow; i < boxStartRow+3; i++ {
		for j := boxStartCol; j < boxStartCol+3; j++ {
			if i == row || j == col {
				continue
			}
			if b[i][j] == val {
				return false
			}
		}
	}

	return true
}

func (b *Board) IsValid() bool {
	for i, row := range b {
		for j, val := range row {
			if !b.IsValidPlacement(i, j, val) {
				return false
			}
		}
	}

	return true
}

func (b *Board) IsSolved() bool {
	for i, row := range b {
		for j, val := range row {
			if val == 0 || !b.IsValidPlacement(i, j, val) {
				return false
			}
		}
	}

	return true
}

func (b *Board) Copy() *Board {
	c := *b
	return &c
}
