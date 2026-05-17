package sudoku

import (
	"testing"
)

func TestSolvePuzzle(t *testing.T) {
	board := &Board{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	solved := &Board{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	success := Solve(board)

	if !success {
		t.Errorf("Expected puzzle to be solved, but it was not")
	}
	for i := range 9 {
		for j := range 9 {
			if board.Get(i, j) != solved.Get(i, j) {
				t.Errorf("Expected cell (%d, %d) to be %d, got %d", i, j, solved.Get(i, j), board.Get(i, j))
			}
		}
	}
}

func TestSolveSolved(t *testing.T) {
	solved := &Board{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	success := Solve(solved)

	if !success {
		t.Errorf("Expected already solved puzzle to be recognized as solved, but it was not")
	}
}

func TestSolveUnsolvable(t *testing.T) {
	unsolvable := &Board{
		{5, 5, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	success := Solve(unsolvable)

	if success {
		t.Errorf("Expected unsolvable puzzle to not be solved, but it was")
	}
}

func TestSolveEmpty(t *testing.T) {
	board := NewBoard()

	success := Solve(board)

	if !success {
		t.Errorf("Expected empty puzzle to be solved, but it was not")
	}
}

func TestSolveStepped(t *testing.T) {
	board := &Board{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	copy := board.Copy()
	solved := &Board{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	ch := SolveStepped(board)

	doneCount := 0
	for step := range ch {
		if step.Action == Place {
			if !copy.IsValidPlacement(step.Row, step.Col, step.Value) {
				t.Fatalf("invalid place: row=%d col=%d val=%d", step.Row, step.Col, step.Value)
			}
			copy.Set(step.Row, step.Col, step.Value)
		} else {
			copy.Set(step.Row, step.Col, 0)
		}

		if step.Action == Done {
			doneCount++
		}
	}
	if doneCount != 1 {
		t.Errorf("Expected exactly one Done action, got %d", doneCount)
	}
	if !board.IsSolved() {
		t.Errorf("Expected puzzle to be solved after stepping through it, but it was not")
	}
	for i := range 9 {
		for j := range 9 {
			if board.Get(i, j) != solved.Get(i, j) {
				t.Errorf("Expected cell (%d, %d) to be %d, got %d", i, j, solved.Get(i, j), board.Get(i, j))
			}
		}
	}
}
