package sudoku

import (
	"errors"
	"math/rand/v2"
)

func countSolutions(b *Board, count *int, limit int) {
	emptyRow, emptyCol := b.findFirstEmptyCell()
	if emptyRow == -1 || emptyCol == -1 {
		if b.IsSolved() {
			*count++
		}
		return
	}

	for val := 1; val <= 9; val++ {
		stepValid := b.checkCellVal(emptyRow, emptyCol, val)
		if !stepValid {
			continue
		}
		b.Set(emptyRow, emptyCol, val)

		countSolutions(b, count, limit)

		b.Set(emptyRow, emptyCol, 0)
		if *count >= limit {
			return
		}
	}
}

func randomFill(b *Board) bool {
	emptyRow, emptyCol := b.findFirstEmptyCell()
	if emptyRow == -1 || emptyCol == -1 {
		return b.IsSolved()
	}

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	for _, num := range nums {
		if b.IsValidPlacement(emptyRow, emptyCol, num) {
			b.Set(emptyRow, emptyCol, num)
			if randomFill(b) {
				return true
			}
			b.Set(emptyRow, emptyCol, 0)
		}
	}

	return false
}

func digHoles(b *Board, amount int) {
	coordinates := make([][2]int, 0, 81)
	for i, row := range b {
		for j := range row {
			coordinates = append(coordinates, [2]int{i, j})
		}
	}
	rand.Shuffle(len(coordinates), func(i, j int) {
		coordinates[i], coordinates[j] = coordinates[j], coordinates[i]
	})

	holes := 0
	for _, coord := range coordinates {
		row, col := coord[0], coord[1]
		original := b.Get(row, col)
		b.Set(row, col, 0)

		solutions := 0
		countSolutions(b, &solutions, 2)
		if solutions == 1 {
			holes++
			if holes >= amount {
				break
			}
		} else {
			b.Set(row, col, original)
		}
	}
}

func Generate() (*Board, error) {
	board := &Board{}

	filled := randomFill(board)
	if !filled {
		return nil, errors.New("failed to generate a complete board")
	}

	digHoles(board, 40)

	return board, nil
}
