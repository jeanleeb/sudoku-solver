package handlers

import (
	"jeanleeb/sudoku-solver/internal/views"
	"jeanleeb/sudoku-solver/sudoku"
	"net/http"
)

func (s *Service) Home(w http.ResponseWriter, r *http.Request) {
	if s.current == nil || s.original == nil {
		puzzle, _ := sudoku.Generate()
		s.original = puzzle
		copy := puzzle.Copy()
		s.current = copy
	}
	component := views.BoardView(s.original, s.current, s.errors)
	views.Layout("Sudoku", component).Render(r.Context(), w)
}
