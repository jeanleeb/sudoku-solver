package sudoku

import "testing"

func TestGetSet(t *testing.T) {
	board := NewBoard()

	board.Set(0, 0, 5)
	board.Set(8, 8, 9)

	if board.Get(0, 0) != 5 {
		t.Errorf("Expected Get(0, 0) to return 5, got %d", board.Get(0, 0))
	}

	if board.Get(8, 8) != 9 {
		t.Errorf("Expected Get(8, 8) to return 9, got %d", board.Get(8, 8))
	}

	if board.Get(4, 4) != 0 {
		t.Errorf("Expected Get(4, 4) to return 0 for an empty cell, got %d", board.Get(4, 4))
	}
}

func TestIsValidPlacement(t *testing.T) {
	board := NewBoard()

	board.Set(0, 1, 5)

	cases := []struct {
		name     string
		input    [3]int
		expected bool
	}{
		{
			name:     "box conflict",
			input:    [3]int{2, 2, 5},
			expected: false,
		},
		{
			name:     "adjacent box, existing value",
			input:    [3]int{1, 4, 5},
			expected: true,
		},
		{
			name:     "same box, different value",
			input:    [3]int{1, 1, 7},
			expected: true,
		},
		{
			name:     "same row, different value",
			input:    [3]int{0, 4, 7},
			expected: true,
		},
		{
			name:     "same column, different value",
			input:    [3]int{2, 1, 7},
			expected: true,
		},
		{
			name:     "same row, existing value",
			input:    [3]int{0, 4, 5},
			expected: false,
		},
		{
			name:     "same column, existing value",
			input:    [3]int{2, 1, 5},
			expected: false,
		},
		{
			name:     "already placed value",
			input:    [3]int{0, 1, 5},
			expected: true,
		},
	}

	for _, c := range cases {
		actual := board.IsValidPlacement(c.input[0], c.input[1], c.input[2])
		if actual != c.expected {
			t.Errorf("Test case '%s' failed: expected %v, got %v", c.name, c.expected, actual)
		}
	}
}

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][3]int
		expected bool
	}{
		{
			name: "invalid board with box conflict",
			input: [][3]int{
				{0, 0, 5},
				{1, 1, 5},
				{2, 2, 7},
			},
			expected: false,
		},
		{
			name: "invalid board with row conflict",
			input: [][3]int{
				{0, 0, 5},
				{0, 1, 5},
			},
			expected: false,
		},
		{
			name: "invalid board with column conflict",
			input: [][3]int{
				{0, 0, 5},
				{1, 0, 5},
			},
			expected: false,
		},
		{
			name: "valid board",
			input: [][3]int{
				{8, 4, 5},
				{1, 6, 3},
			},
			expected: true,
		},
	}

	for _, c := range testCases {
		board := NewBoard()
		for _, placement := range c.input {
			board.Set(placement[0], placement[1], placement[2])
		}
		actual := board.IsValid()
		if actual != c.expected {
			t.Errorf("Test case '%s' failed: expected %v, got %v", c.name, c.expected, actual)
		}
	}
}

func TestIsSolved(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][3]int
		expected bool
	}{
		{
			name:     "empty board",
			input:    [][3]int{},
			expected: false,
		},
		{
			name: "incomplete board",
			input: [][3]int{
				{0, 0, 5},
				{1, 1, 5},
			},
			expected: false,
		},
		{
			name: "complete and invalid board",
			input: [][3]int{
				{0, 0, 8}, {0, 1, 8}, {0, 2, 5}, {0, 3, 2}, {0, 4, 6}, {0, 5, 9}, {0, 6, 7}, {0, 7, 1}, {0, 8, 3},
				{1, 0, 1}, {1, 1, 6}, {1, 2, 3}, {1, 3, 7}, {1, 4, 9}, {1, 5, 8}, {1, 6, 4}, {1, 7, 5}, {1, 8, 2},
				{2, 0, 7}, {2, 1, 9}, {2, 2, 2}, {2, 3, 4}, {2, 4, 5}, {2, 5, 3}, {2, 6, 8}, {2, 7, 6}, {2, 8, 1},
				{3, 0, 9}, {3, 1, 8}, {3, 2, 6}, {3, 3, 5}, {3, 4, 7}, {3, 5, 4}, {3, 6, 2}, {3, 7, 3}, {3, 8, 5},
				{4, 0, 4}, {4, 1, 5}, {4, 2, 7}, {4, 3, 6}, {4, 4, 8}, {4, 5, 2}, {4, 6, 9}, {4, 7, 1}, {4, 8, 8},
				{5, 0, 2}, {5, 1, 3}, {5, 2, 8}, {5, 3, 9}, {5, 4, 1}, {5, 5, 7}, {5, 6, 6}, {5, 7, 4}, {5, 8, 9},
				{6, 0, 6}, {6, 1, 7}, {6, 2, 4}, {6, 3, 8}, {6, 4, 3}, {6, 5, 5}, {6, 6, 1}, {6, 7, 9}, {6, 8, 7},
				{7, 0, 5}, {7, 1, 2}, {7, 2, 9}, {7, 3, 1}, {7, 4, 4}, {7, 5, 6}, {7, 6, 3}, {7, 7, 8}, {7, 8, 4},
				{8, 0, 3}, {8, 1, 1}, {8, 2, 8}, {8, 3, 3}, {8, 4, 2}, {8, 5, 7}, {8, 6, 5}, {8, 7, 4}, {8, 8, 6},
			},
			expected: false,
		},
		{
			name: "complete and valid board",
			input: [][3]int{
				{0, 0, 5}, {0, 1, 3}, {0, 2, 4}, {0, 3, 6}, {0, 4, 7}, {0, 5, 8}, {0, 6, 9}, {0, 7, 1}, {0, 8, 2},
				{1, 0, 6}, {1, 1, 7}, {1, 2, 2}, {1, 3, 1}, {1, 4, 9}, {1, 5, 5}, {1, 6, 3}, {1, 7, 4}, {1, 8, 8},
				{2, 0, 1}, {2, 1, 9}, {2, 2, 8}, {2, 3, 3}, {2, 4, 4}, {2, 5, 2}, {2, 6, 5}, {2, 7, 6}, {2, 8, 7},
				{3, 0, 8}, {3, 1, 5}, {3, 2, 9}, {3, 3, 7}, {3, 4, 6}, {3, 5, 1}, {3, 6, 4}, {3, 7, 2}, {3, 8, 3},
				{4, 0, 4}, {4, 1, 2}, {4, 2, 6}, {4, 3, 8}, {4, 4, 5}, {4, 5, 3}, {4, 6, 7}, {4, 7, 9}, {4, 8, 1},
				{5, 0, 7}, {5, 1, 1}, {5, 2, 3}, {5, 3, 9}, {5, 4, 2}, {5, 5, 4}, {5, 6, 8}, {5, 7, 5}, {5, 8, 6},
				{6, 0, 9}, {6, 1, 6}, {6, 2, 1}, {6, 3, 5}, {6, 4, 3}, {6, 5, 7}, {6, 6, 2}, {6, 7, 8}, {6, 8, 4},
				{7, 0, 2}, {7, 1, 8}, {7, 2, 7}, {7, 3, 4}, {7, 4, 1}, {7, 5, 9}, {7, 6, 6}, {7, 7, 3}, {7, 8, 5},
				{8, 0, 3}, {8, 1, 4}, {8, 2, 5}, {8, 3, 2}, {8, 4, 8}, {8, 5, 6}, {8, 6, 1}, {8, 7, 7}, {8, 8, 9},
			},
			expected: true,
		},
	}

	for _, c := range testCases {
		board := NewBoard()
		for _, placement := range c.input {
			board.Set(placement[0], placement[1], placement[2])
		}
		actual := board.IsSolved()
		if actual != c.expected {
			t.Errorf("Test case '%s' failed: expected %v, got %v", c.name, c.expected, actual)
		}
	}
}

func TestCopy(t *testing.T) {
	board := NewBoard()
	board.Set(0, 0, 5)
	board.Set(8, 8, 9)

	copy := board.Copy()

	if copy.Get(0, 0) != 5 {
		t.Errorf("Expected copied board Get(0, 0) to return 5, got %d", copy.Get(0, 0))
	}
	if copy.Get(8, 8) != 9 {
		t.Errorf("Expected copied board Get(8, 8) to return 9, got %d", copy.Get(8, 8))
	}
	if copy.Get(4, 4) != 0 {
		t.Errorf("Expected copied board Get(4, 4) to return 0 for an empty cell, got %d", copy.Get(4, 4))
	}

	copy.Set(0, 0, 7)
	copy.Set(8, 8, 1)

	if board.Get(0, 0) != 5 {
		t.Errorf("Expected original board Get(0, 0) to remain 5 after modifying copy, got %d", board.Get(0, 0))
	}
	if board.Get(8, 8) != 9 {
		t.Errorf("Expected original board Get(8, 8) to remain 9 after modifying copy, got %d", board.Get(8, 8))
	}
}
