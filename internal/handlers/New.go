package handlers

import (
	"jeanleeb/sudoku-solver/sudoku"
	"net/http"
)

func (s *Service) New(w http.ResponseWriter, r *http.Request) {
	puzzle, _ := sudoku.Generate()

	s.original = puzzle
	copy := puzzle.Copy()
	s.current = copy
	s.errors = [9][9]bool{}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
