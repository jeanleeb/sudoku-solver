package sudoku

import (
	"testing"
)

func TestCountSolutions(t *testing.T) {
	cases := []struct {
		name     string
		board    *Board
		expected int
	}{
		{
			name: "Unique solution",
			board: &Board{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			expected: 1,
		},
		{
			name: "Multiple solutions with limit",
			board: &Board{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{4, 5, 6, 7, 8, 9, 1, 2, 3},
				{7, 8, 9, 1, 2, 3, 4, 5, 6},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: 2,
		},
		{
			name: "No solutions",
			board: &Board{
				{5, 5, 5, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9},
			},
			expected: 0,
		},
	}

	for _, c := range cases {
		count := 0
		copy := c.board.Copy()
		countSolutions(copy, &count, 2)

		if count != c.expected {
			t.Errorf("Test %s: expected %d solutions, got %d", c.name, c.expected, count)
		}
		for i, row := range copy {
			for j, val := range row {
				originalVal := c.board.Get(i, j)
				if val != originalVal {
					t.Errorf("Test %s: expected cell (%d, %d) to be %d, got %d", c.name, i, j, originalVal, val)
				}
			}
		}
	}
}

func TestRandomFill(t *testing.T) {
	board := &Board{}
	success := randomFill(board)

	if !success {
		t.Errorf("Expected randomFill to succeed, but it failed")
	}
	if !board.IsSolved() {
		t.Errorf("Expected board to be solved after randomFill, but it was not")
	}
	for _, row := range board {
		for _, val := range row {
			if val < 1 || val > 9 {
				t.Errorf("Expected all values to be between 1 and 9, but got %d", val)
			}
		}
	}
}

func TestDigHoles(t *testing.T) {
	board := &Board{}
	randomFill(board)

	digHoles(board, 50)

	holes := 0
	for i := range 9 {
		for j := range 9 {
			if board.Get(i, j) == 0 {
				holes++
			}
		}
	}
	if holes < 30 {
		t.Errorf("Expected at least 30 holes, but got %d", holes)
	}
}

func TestGenerate(t *testing.T) {
	board, err := Generate()

	if err != nil {
		t.Errorf("Expected Generate to succeed, but it failed with error: %v", err)
	}
	if !board.IsValid() {
		t.Errorf("Expected generated board to be valid, but it was not")
	}
	solutions := 0
	countSolutions(board.Copy(), &solutions, 2)
	if solutions != 1 {
		t.Errorf("Expected generated board to have a unique solution, but it has %d solutions", solutions)
	}
	if board.IsSolved() {
		t.Errorf("Expected generated board to have holes, but it is already solved")
	}
}
