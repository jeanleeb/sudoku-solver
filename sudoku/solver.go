package sudoku

func (b *Board) findFirstEmptyCell() (int, int) {
	for i, row := range b {
		for j, val := range row {
			if val == 0 {
				return i, j
			}
		}
	}
	return -1, -1 // No empty cell found
}

func (b *Board) checkCellVal(row, col, val int) bool {
	if b.Get(row, col) != 0 {
		return false
	}
	if !b.IsValidPlacement(row, col, val) {
		return false
	}

	b.Set(row, col, val)
	return true
}

func Solve(b *Board) bool {
	emptyRow, emptyCol := b.findFirstEmptyCell()
	if emptyRow == -1 || emptyCol == -1 {
		return b.IsSolved()
	}

	for val := 1; val <= 9; val++ {
		stepRes := b.checkCellVal(emptyRow, emptyCol, val)
		if !stepRes {
			continue
		}
		if Solve(b) {
			return true
		}
		b.Set(emptyRow, emptyCol, 0)
	}

	return false
}

type Action string

const (
	Place     Action = "place"
	Backtrack Action = "backtrack"
	Done      Action = "done"
)

type Step struct {
	Row    int
	Col    int
	Value  int
	Action Action
}

func solveStepped(b *Board, ch chan<- Step) bool {
	emptyRow, emptyCol := b.findFirstEmptyCell()
	if emptyRow == -1 || emptyCol == -1 {
		return b.IsSolved()
	}

	for val := 1; val <= 9; val++ {
		stepRes := b.checkCellVal(emptyRow, emptyCol, val)
		if !stepRes {
			continue
		}
		ch <- Step{
			Row:    emptyRow,
			Col:    emptyCol,
			Value:  val,
			Action: Place,
		}

		if solveStepped(b, ch) {
			return true
		}

		ch <- Step{
			Row:    emptyRow,
			Col:    emptyCol,
			Value:  val,
			Action: Backtrack,
		}
		b.Set(emptyRow, emptyCol, 0)
	}

	return false
}

func SolveStepped(b *Board) <-chan Step {
	ch := make(chan Step, 64)

	go func() {
		defer close(ch)
		solveStepped(b, ch)
		ch <- Step{
			Action: Done,
		}
	}()

	return ch
}
