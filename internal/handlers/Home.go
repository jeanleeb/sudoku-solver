package handlers

import (
	"jeanleeb/sudoku-solver/internal/views"
	"jeanleeb/sudoku-solver/sudoku"
	"net/http"
)

func (s *Service) Home(w http.ResponseWriter, r *http.Request) {
	if s.current == nil {
		puzzle, _ := sudoku.Generate()
		s.current = puzzle
	}
	component := views.BoardView(s.current, [9][9]bool{})
	views.Layout("Sudoku", component).Render(r.Context(), w)
}
